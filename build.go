package main

import (
	"fmt"
	"bytes"
	"math"
	"github.com/PuerkitoBio/goquery"
	"os"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"io/ioutil"
	"strconv"
)

type Range struct {
	Min int
	Max int
}

func buildAllHTMLS() {

	htmLFiles := getHtmlFiles()

	for i, theRange := range getRanges(len(htmLFiles)) {
		var pdfFiles []HtmlFile
		for i := theRange.Min; i <= theRange.Max; i++ {
			pdfFiles = append(pdfFiles, htmLFiles[i-1])
		}
		// PrettyPrint(pdfFiles)
		// pp("")
		createHTML(pdfFiles, theRange, i)
	}

}
func createHTML(files []HtmlFile, theRange Range, fileNo int) {

	fCount := len(files)
	firstHtmlFile := files[0]

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(firstHtmlFile.Content)))
	pErr("goquery.NewDocumentFromReader", err)

	head := doc.Find("head")

	head.AppendHtml("<style>" + ConstPageBreakCss + "</style>")

	if fileExists(cfg.CustomCssFile) {
		head.AppendHtml(`<style>` + getFileStr(cfg.CustomCssFile) + `</style>`)
	}


	for i := 1; i < fCount; i++ {
		articleParent := doc.Find(cfg.ArticleParentElement)
		articleParent.AppendHtml(`<hr/>`)
		articleParent.AppendHtml(getContent(files[i]))
	}

	doc.Find(cfg.ArticleTitleClass).Each(func(i int, s *goquery.Selection) {
		s.PrependHtml("[" + strconv.Itoa(theRange.Min+i) + "] ")
	})

	docHtmlStr, err := doc.Selection.Html()

	if err != nil {
		pp("Error doc.Selection.Html:" + err.Error())
	}

	htmlFilePath := fmt.Sprintf(combinedHtmlDir+"/%d-%d_"+cfg.Domain+".html", theRange.Min, theRange.Max)

	osFile, err := os.Create(htmlFilePath)
	if err != nil {
		pp("Error os.Create(htmlFilePath): " + err.Error())
	}

	p(fmt.Sprintf("%d: Generated Combined HTML File: "+htmlFilePath, fileNo+1))

	osFile.WriteString(docHtmlStr)

	osFile.Close()

	if cfg.GeneratePDF {
		htmlToPDF(htmlFilePath,
			fmt.Sprintf(pdfDir+"/%d-%d_"+cfg.Domain+".pdf", theRange.Min, theRange.Max), fileNo,
		)
	}
}

func htmlToPDF(htmlFilePath string, pdfFilePath string, i int) {

	p(fmt.Sprintf("%d: Creating PDF File!", i+1))

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		pp("Error wkhtmltopdf.NewPDFGenerator: " + err.Error())
	}

	htmlfile, err := ioutil.ReadFile(htmlFilePath)
	if err != nil {
		pp("Error ioutil.ReadFile(htmlFilePath): " + err.Error())
	}

	pdfg.PageSize.Set(cfg.PdfPageSize)
	pdfg.MarginLeft.Set(uint(cfg.PdfMarginTop))
	pdfg.MarginRight.Set(uint(cfg.PdfMarginLeft))
	pdfg.MarginTop.Set(uint(cfg.PdfMarginRight))
	pdfg.MarginBottom.Set(uint(cfg.PdfMarginBottom))
	pdfg.Orientation.Set(cfg.PdfOrientation)

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(htmlfile)))

	err = pdfg.Create()
	if err != nil {
		p("Error while pdfg.Create: " + err.Error())
	}

	err = pdfg.WriteFile(pdfFilePath)
	if err != nil {
		pp("Error pdfg.WriteFile: " + err.Error())
	}

	p(fmt.Sprintf("%d: Generated PDF size %vkB: %v\n", i+1, len(pdfg.Bytes())/1024, pdfFilePath))

	if pdfg.Buffer().Len() != len(pdfg.Bytes()) {
		fmt.Println("Buffersize not equal: " + pdfFilePath)
	}
}

func getRanges(totalHtmlCount int) []Range {

	var ranges []Range

	// ps("totalHtmlFiles: " + strconv.Itoa(totalHtmlCount))
	totalPdfCount := int(math.Floor(float64(totalHtmlCount) / float64(cfg.ArticlePerPDF)))
	// pm("totalPdfCount: " + strconv.Itoa(totalPdfCount))
	lastMax := 0

	for i := 0; i < totalPdfCount; i++ {

		min := lastMax + 1
		max := (i + 1) * cfg.ArticlePerPDF

		if max > totalHtmlCount {
			max = totalHtmlCount
		}

		lastMax = max

		ranges = append(ranges, Range{Min: min, Max: max})

	}

	if totalHtmlCount > lastMax {
		ranges = append(ranges, Range{Min: lastMax + 1, Max: totalHtmlCount})
	}

	// PrettyPrint(ranges)

	return ranges

}
