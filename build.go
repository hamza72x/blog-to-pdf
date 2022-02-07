package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	hel "github.com/hamza72x/go-helper"
	"github.com/yosssi/gohtml"
)

func build() {

	allHTMLFiles := getHTMLFiles()

	removeContents(combinedHTMLDir)

	var wg sync.WaitGroup
	var c = make(chan int, thread)

	for i, pdfile := range getPdfiles(allHTMLFiles) {
		wg.Add(1)
		go func(pdfile xPdfile, i int) {
			c <- i
			buildCombinedHTMLAndGeneratePDF(pdfile)
			wg.Done()
			<-c
		}(pdfile, i)
	}

	wg.Wait()
	close(c)
}
func getPdfiles(allHTMLFiles []xHTMLFile) []xPdfile {
	var pdfiles []xPdfile

	for i, theRange := range getRanges(len(allHTMLFiles)) {

		var htmls []xHTMLFile

		for i := theRange.Start; i <= theRange.End; i++ {
			htmls = append(htmls, allHTMLFiles[i-1])
		}

		pdfiles = append(pdfiles,
			xPdfile{i + 1, htmls, theRange},
		)
	}
	return pdfiles
}

func buildCombinedHTMLAndGeneratePDF(pdfile xPdfile) {

	// one pdf/doc will have multiple html file
	creditHTML := hel.FileStrMust("credit.html")
	htmls := pdfile.HTMLFiles
	theRange := pdfile.TheRange
	serial := pdfile.Serial

	htmlTemplate, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(htmlTemplate)))
	hel.PlP("[1] goquery.NewDocumentFromReader", err)

	htmlContainer := htmlTemplate.Find("div.the-tool-container")
	htmlHead := htmlTemplate.Find("head")

	if hel.FileExists(cfg.CustomCSSFile) {
		htmlHead.AppendHtml(`<style>` + hel.FileStrMust(cfg.CustomCSSFile) + `</style>`)
	}

	htmlContainer.AppendHtml(strings.ReplaceAll(creditHTML,
		"title_placeholder",
		fmt.Sprintf("%d-%d_"+cfg.PdfFileName, theRange.Start, theRange.End)),
	)
	// combine html files
	// according to range
	regexHTMLComment := regexp.MustCompile(`<.*?>`)

	for i := 0; i < len(htmls); i++ {

		content, err := getContentHTML(htmls[i])

		if err != nil {
			hel.Pl("🥺 getContentHTML", err, "SKIPPING")
			continue
		}

		if len(content) == 0 {
			hel.Pl("🥺 Empty content of", htmls[i], "SKIPPING")
			continue
		}

		contentArr := strings.Fields(strings.TrimSpace(regexHTMLComment.ReplaceAllString(content, "")))
		estTime := int(math.Ceil(float64(len(contentArr) / 170)))

		htmlContainer.AppendHtml(
			fmt.Sprintf(
				`<article class="general-article"><div><p class="text-center">%d MIN READ</p></div>%s</article>`,
				estTime,
				content,
			),
		)
	}

	// set [i] in title
	htmlContainer.Find(cfg.ArticleTitleClass).Each(func(i int, s *goquery.Selection) {
		if cfg.AppendAutoArticleNumberInTitle {
			s.PrependHtml("[" + strconv.Itoa(theRange.Start+i) + "] ")
			s.AddClass("text-center")
		}
		if cfg.AppendURLInTitle {
			s.AppendHtml(
				fmt.Sprintf(
					"<br/><a class=\"article-origin-link\" href=\"%s\">%s</a>",
					htmls[i].RemoteURL,
					`Article Link`,
				),
			)
		}
	})

	htmlContainer.AppendHtml(
		strings.ReplaceAll(
			strings.ReplaceAll(creditHTML, `the-page-break-class`, ""),
			"title_placeholder",
			fmt.Sprintf("%d-%d_"+cfg.PdfFileName, theRange.Start, theRange.End),
		),
	)

	combinedHTMLStr, err := htmlTemplate.Html()
	hel.PlP("doc.Selection.Html", err)

	combinedHTMLStr = gohtml.Format(combinedHTMLStr)

	combinedHTMLFilePath := fmt.Sprintf(
		combinedHTMLDir+"/%d-%d_"+cfg.Domain+".html",
		theRange.Start,
		theRange.End,
	)

	osFile, err := os.Create(combinedHTMLFilePath)
	hel.PlP("os.Create(htmlFilePath)", err)

	hel.Pl(fmt.Sprintf("🔪 %d: Generated Combined HTML File: "+combinedHTMLFilePath, serial))

	osFile.WriteString(combinedHTMLStr)

	osFile.Close()

	if cfg.GeneratePDF {

		var pdfFilePath = fmt.Sprintf(cfg.PdfOutputDirPath+"/%d-%d_"+cfg.PdfFileName+".pdf", theRange.Start, theRange.End)

		if cfg.SkipPDFCreationIfExistsAlready && hel.FileExists(pdfFilePath) {
			hel.Pl(fmt.Sprintf("%d: [SkipPDFCreationIfExistsAlready] Already exists!", serial))
		} else {
			htmlToPDF(combinedHTMLFilePath, pdfFilePath, serial)
		}

	} else {
		hel.Pl("Skipping pdf generation as it's disabled in `" + cfgFilePath + "`")
	}
}

func htmlToPDF(htmlFilePath string, pdfFilePath string, serial int) {

	hel.Pl(fmt.Sprintf("🛠 %d: Creating PDF File!", serial))

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	hel.PlP("wkhtmltopdf.NewPDFGenerator", err)

	htmlfile, err := ioutil.ReadFile(htmlFilePath)
	hel.PlP("ioutil.ReadFile(htmlFilePath)", err)

	pdfg.PageSize.Set(cfg.PdfPageSize)
	pdfg.MarginLeft.Set(uint(cfg.PdfMarginTop))
	pdfg.MarginRight.Set(uint(cfg.PdfMarginLeft))
	pdfg.MarginTop.Set(uint(cfg.PdfMarginRight))
	pdfg.MarginBottom.Set(uint(cfg.PdfMarginBottom))
	pdfg.Orientation.Set(cfg.PdfOrientation)

	// pdfg.AddPage(wkhtmltopdf.NewPage(creditHtml))
	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(htmlfile)))
	// pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewBufferString(creditHtml)))

	err = pdfg.Create()
	hel.PlP("pdfg.Create", err)

	err = pdfg.WriteFile(pdfFilePath)
	hel.PlP("pdfg.WriteFile", err)

	hel.Pl(fmt.Sprintf("✅ %d: Generated PDF size %vkB: %v", serial, len(pdfg.Bytes())/1024, pdfFilePath))

	if pdfg.Buffer().Len() != len(pdfg.Bytes()) {
		fmt.Println("Buffersize not equal: " + pdfFilePath)
	}
}

func getRanges(totalHTMLCount int) []xRange {

	var ranges []xRange

	totalPdfCount := int(math.Floor(float64(totalHTMLCount) / float64(cfg.ArticlePerPDF)))
	lastEnd := 0

	for i := 0; i < totalPdfCount; i++ {

		start := lastEnd + 1
		end := (i + 1) * cfg.ArticlePerPDF

		if end > totalHTMLCount {
			end = totalHTMLCount
		}

		lastEnd = end

		ranges = append(ranges, xRange{Start: start, End: end})

	}

	if totalHTMLCount > lastEnd {
		ranges = append(ranges, xRange{Start: lastEnd + 1, End: totalHTMLCount})
	}

	return ranges

}

// removes all contens in a given directory
func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
