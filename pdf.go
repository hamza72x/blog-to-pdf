package main

import (
	"fmt"
	"bytes"
	"math"
	"github.com/PuerkitoBio/goquery"
	"os"
)

type Range struct {
	iMin int
	iMax int
}

func generateAllPdf() {

	files := getHtmlFiles(forceFetchHtml, forceSiteMapFetch)

	totalFiles := len(files)

	var articleRanges []Range

	pdfs := int(math.Floor(float64(totalFiles) / float64(ArticlePerPDF)))
	lastIMax := 0

	for i := 0; i < pdfs; i++ {
		iMax := (i + 1) * ArticlePerPDF
		if iMax == totalFiles {
			iMax = iMax - 1
		}
		lastIMax = iMax
		articleRanges = append(articleRanges, Range{iMin: i * ArticlePerPDF, iMax: iMax})
	}

	if totalFiles-1 > lastIMax {
		articleRanges = append(articleRanges, Range{iMin: lastIMax + 1, iMax: totalFiles - 1})
	}

	// fmt.Printf("articleRanges: %+v", articleRanges)

	for _, theRange := range articleRanges {
		var pdfFiles []HtmlFile
		for i := theRange.iMin; i <= theRange.iMax; i++ {
			pdfFiles = append(pdfFiles, files[i])
		}
		createPDF(pdfFiles, theRange)
	}

}
func createPDF(files []HtmlFile, theRange Range) {

	fCount := len(files)
	firstHtmlFile := files[0]

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(firstHtmlFile.Content)))
	check(err)

	for i := 1; i < fCount; i++ {
		doc.Find("#content").AppendHtml(getContent(files[i]))
	}

	docHtmlStr, err := doc.Selection.Html()
	check(err)

	filename := fmt.Sprintf(buildDir+"/[%d-%d] "+DOMAIN+".html", theRange.iMin, theRange.iMax)

	osFile, err := os.Create(filename)

	check(err)

	fmt.Println("Generated: " + filename)

	osFile.WriteString(docHtmlStr)

	osFile.Close()
}
