package main

import (
	"fmt"
	"bytes"
	"math"
	"github.com/PuerkitoBio/goquery"
	"os"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"io/ioutil"
)

type Range struct {
	iMin int
	iMax int
}

func buildAllHTMLS() {

	var articlePerPDF = iniGetInt("article_per_pdf")
	files := getHtmlFiles()

	totalFiles := len(files)

	var articleRanges []Range

	totalPdfCount := int(math.Floor(float64(totalFiles) / float64(articlePerPDF)))
	lastIMax := 0

	for i := 0; i < totalPdfCount; i++ {

		iMax := (i + 1) * articlePerPDF
		iMin := i * articlePerPDF

		if iMax == totalFiles {
			iMax = iMax - 1
		}

		if iMin != 0 {
			iMin += 1
		}

		lastIMax = iMax

		articleRanges = append(articleRanges, Range{iMin: iMin, iMax: iMax})
	}

	if totalFiles-1 > lastIMax {
		articleRanges = append(articleRanges, Range{iMin: lastIMax + 1, iMax: totalFiles - 1})
	}

	// fmt.Printf("articleRanges: %+v", articleRanges)
	// panic("END")

	for i, theRange := range articleRanges {
		var pdfFiles []HtmlFile
		for i := theRange.iMin; i <= theRange.iMax; i++ {
			pdfFiles = append(pdfFiles, files[i])
		}
		createHTML(pdfFiles, theRange, i)
	}

}
func createHTML(files []HtmlFile, theRange Range, fileNo int) {

	fCount := len(files)
	firstHtmlFile := files[0]

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(firstHtmlFile.Content)))
	if err != nil {
		pp("Error goquery.NewDocumentFromReader:" + err.Error())
	}

	for i := 1; i < fCount; i++ {
		doc.Find(articleParentElement).AppendHtml(getContent(files[i]))
	}

	docHtmlStr, err := doc.Selection.Html()
	if err != nil {
		pp("Error doc.Selection.Html:" + err.Error())
	}

	htmlFilePath := fmt.Sprintf(combinedHtmlDir+"/%d-%d_"+siteDomain+".html", theRange.iMin+1, theRange.iMax+1)

	osFile, err := os.Create(htmlFilePath)
	if err != nil {
		pp("Error os.Create(htmlFilePath): " + err.Error())
	}

	p(fmt.Sprintf("%d: Generated Combined HTML File: "+htmlFilePath, fileNo+1))

	osFile.WriteString(docHtmlStr)

	osFile.Close()

	if generatePDF {
		htmlToPDF(htmlFilePath,
			fmt.Sprintf(pdfDir+"/%d-%d_"+siteDomain+".pdf", theRange.iMin+1, theRange.iMax+1), fileNo,
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

	pdfg.PageSize.Set(pdfPageSize)
	pdfg.MarginLeft.Set(uint(pdfMarginTop))
	pdfg.MarginRight.Set(uint(pdfMarginLeft))
	pdfg.MarginTop.Set(uint(pdfMarginRight))
	pdfg.MarginBottom.Set(uint(pdfMarginBottom))

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
