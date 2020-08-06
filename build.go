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

	"github.com/PuerkitoBio/goquery"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	hel "github.com/thejini3/go-helper"
	"github.com/yosssi/gohtml"
)

func build() {

	allHTMLFiles := getHTMLFiles()

	removeContents(combinedHTMLDir)

	for _, pdfile := range getPdfiles(allHTMLFiles) {
		buildCombinedHTMLAndGeneratePDF(pdfile)
	}

}
func getPdfiles(allHTMLFiles []xHTMLFile) (outFiles []xPdfile) {

	for i, theRange := range getRanges(len(allHTMLFiles)) {

		var htmls []xHTMLFile

		for i := theRange.Min; i <= theRange.Max; i++ {
			htmls = append(htmls, allHTMLFiles[i-1])
		}

		outFiles = append(outFiles,
			xPdfile{htmls, theRange, i},
		)
	}
	return
}

func buildCombinedHTMLAndGeneratePDF(pdfile xPdfile) {

	// one pdf/doc will have multiple html file
	htmls := pdfile.HTMLFiles
	theRange := pdfile.TheRange
	fileNo := pdfile.FileNo

	htmlTemplate, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(htmlTemplate)))
	hel.PErr("[1] goquery.NewDocumentFromReader", err)

	htmlContainer := htmlTemplate.Find("div.the-tool-container")
	htmlHead := htmlTemplate.Find("head")

	if hel.FileExists(cfg.CustomCSSFile) {
		htmlHead.AppendHtml(`<style>` + hel.GetFileStr(cfg.CustomCSSFile) + `</style>`)
	}

	htmlContainer.AppendHtml(strings.ReplaceAll(frontAndBackPage,
		"title_placeholder",
		fmt.Sprintf("%d-%d_"+cfg.PdfFileName, theRange.Min, theRange.Max)),
	)
	// combine html files
	// according to range
	regexHTMLComment := regexp.MustCompile(`<.*?>`)

	for i := 0; i < len(htmls); i++ {

		content := getContentHTML(htmls[i])
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
			s.PrependHtml("[" + strconv.Itoa(theRange.Min+i) + "] ")
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
			strings.ReplaceAll(frontAndBackPage, `the-page-break-class`, ""),
			"title_placeholder",
			fmt.Sprintf("%d-%d_"+cfg.PdfFileName, theRange.Min, theRange.Max),
		),
	)

	combinedHTMLStr, err := htmlTemplate.Html()
	hel.PErr("doc.Selection.Html", err)

	combinedHTMLStr = gohtml.Format(combinedHTMLStr)

	combinedHTMLFilePath := fmt.Sprintf(
		combinedHTMLDir+"/%d-%d_"+cfg.Domain+".html",
		theRange.Min,
		theRange.Max,
	)

	osFile, err := os.Create(combinedHTMLFilePath)
	hel.PErr("os.Create(htmlFilePath)", err)

	hel.PS(fmt.Sprintf("%d: Generated Combined HTML File: "+combinedHTMLFilePath, fileNo+1))

	osFile.WriteString(combinedHTMLStr)

	osFile.Close()

	if cfg.GeneratePDF {

		var pdfFilePath = fmt.Sprintf(cfg.PdfOutputDirPath+"/%d-%d_"+cfg.PdfFileName+".pdf", theRange.Min, theRange.Max)

		if cfg.SkipPDFCreationIfExistsAlready && hel.FileExists(pdfFilePath) {
			hel.PE(fmt.Sprintf("%d: [SkipPDFCreationIfExistsAlready] Already exists!", fileNo))
		} else {
			htmlToPDF(combinedHTMLFilePath, pdfFilePath, fileNo)
		}

	} else {
		hel.PE("Skipping pdf generation as it's disabled in `" + cfgFilePath + "`")
	}
}

func htmlToPDF(htmlFilePath string, pdfFilePath string, i int) {

	hel.PM(fmt.Sprintf("%d: Creating PDF File!", i+1))

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	hel.ErrOSExit("wkhtmltopdf.NewPDFGenerator", err)

	htmlfile, err := ioutil.ReadFile(htmlFilePath)
	hel.ErrOSExit("ioutil.ReadFile(htmlFilePath)", err)

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
	hel.ErrOSExit("pdfg.Create", err)

	err = pdfg.WriteFile(pdfFilePath)
	hel.ErrOSExit("pdfg.WriteFile", err)

	hel.PE(fmt.Sprintf("%d: Generated PDF size %vkB: %v", i+1, len(pdfg.Bytes())/1024, pdfFilePath))

	if pdfg.Buffer().Len() != len(pdfg.Bytes()) {
		fmt.Println("Buffersize not equal: " + pdfFilePath)
	}
}

func getRanges(totalHTMLCount int) []xRange {

	var ranges []xRange

	totalPdfCount := int(math.Floor(float64(totalHTMLCount) / float64(cfg.ArticlePerPDF)))
	lastMax := 0

	for i := 0; i < totalPdfCount; i++ {

		min := lastMax + 1
		max := (i + 1) * cfg.ArticlePerPDF

		if max > totalHTMLCount {
			max = totalHTMLCount
		}

		lastMax = max

		ranges = append(ranges, xRange{Min: min, Max: max})

	}

	if totalHTMLCount > lastMax {
		ranges = append(ranges, xRange{Min: lastMax + 1, Max: totalHTMLCount})
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
