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

	files := getHtmlFiles()

	totalFiles := len(files)

	var articleRanges []Range

	totalPdfCount := int(math.Floor(float64(totalFiles) / float64(ArticlePerPDF)))
	lastIMax := 0

	for i := 0; i < totalPdfCount; i++ {

		iMax := (i + 1) * ArticlePerPDF
		iMin := i * ArticlePerPDF

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
	check(err)

	for i := 1; i < fCount; i++ {
		doc.Find(getArticleWrapper()).AppendHtml(getContent(files[i]))
	}

	docHtmlStr, err := doc.Selection.Html()
	check(err)

	htmlFilePath := fmt.Sprintf(buildDir+"/%d-%d_"+DOMAIN+".html", theRange.iMin + 1, theRange.iMax + 1)

	osFile, err := os.Create(htmlFilePath)
	check(err)

	p(fmt.Sprintf("%d: Generated Combined HTML File: " + htmlFilePath, fileNo + 1))

	osFile.WriteString(docHtmlStr)

	osFile.Close()

	if generatePdf {
		htmlToPDF(htmlFilePath,
			fmt.Sprintf(pdfDir+"/%d-%d_"+DOMAIN+".pdf", theRange.iMin + 1, theRange.iMax + 1), fileNo,
		)
	}
}

func htmlToPDF(htmlFilePath string, pdfFilePath string, i int) {

	p(fmt.Sprintf("%d: Creating PDF File!", i + 1))

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	check(err)

	htmlfile, err := ioutil.ReadFile(htmlFilePath)
	check(err)

	pdfg.PageSize.Set(pdfPageSize)
	pdfg.MarginLeft.Set(uint(pdfMargin))
	pdfg.MarginRight.Set(uint(pdfMargin))
	pdfg.MarginTop.Set(uint(pdfMargin))
	pdfg.MarginBottom.Set(uint(pdfMargin))

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(htmlfile)))

	err = pdfg.Create()
	if err != nil {
		p("Error while pdfg.Create: " + err.Error())
	}

	err = pdfg.WriteFile(pdfFilePath)
	check(err)

	p(fmt.Sprintf("%d: Generated PDF size %vkB: %v\n", i + 1, len(pdfg.Bytes())/1024, pdfFilePath))

	if pdfg.Buffer().Len() != len(pdfg.Bytes()) {
		fmt.Println("Buffersize not equal: " + pdfFilePath)
	}
}


