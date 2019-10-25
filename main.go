package main

import (
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"os"
	"fmt"
	"math"
)

const TempHtml = "build/temp.html"
const DOMAIN = "alorpothe.wordpress.com"
const SITE = "https://" + DOMAIN
const SiteMapURL = "https://" + DOMAIN + "/sitemap.xml"
const UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1"
const ArticlePerPDF = 10

var SiteBasedTags = map[string][]string{
	"alorpothe.wordpress.com": {"div#secondary", "div.menu-search", "nav#nav-single",
		"footer#colophon", "div.widget", "div#fb-root", "div#actionbar",
		"form", "h3#reply-title", "div#jp-post-flair", "div#header-img", ".cs-rating .pd-rating",
		"div.wpcnt", "h3#entry-format", ".rating-star-icon",
	},
}

type Range struct {
	iMin int
	iMax int
}

func main() {

	files := getHtmlFiles(false, false)
	totalFiles := len(files)
	// total files = 200
	fmt.Println("total files: ", totalFiles)

	var aticleRanges []Range

	pdfs := int(math.Floor(float64(totalFiles) / ArticlePerPDF))

	for i := 0; i < pdfs; i++ {
		iMax := (i + 1) * ArticlePerPDF
		if iMax == totalFiles {
			iMax = iMax - 1
		}
		aticleRanges = append(aticleRanges, Range{iMin: i * ArticlePerPDF, iMax: iMax})
	}

	fmt.Printf("aticleRanges: %+v", aticleRanges)

	for _, theRange := range aticleRanges {
		var combineFiles []HtmlFile
		for i := theRange.iMin; i <= theRange.iMax; i++ {
			combineFiles = append(combineFiles, files[i])
		}
		combine(combineFiles, theRange)
	}

}

func combine(files []HtmlFile, theRange Range) {

	var articles string
	fCount := len(files)
	firstHtmlFile := files[0]

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(firstHtmlFile.Content)))
	check(err)

	for i := 1; i < fCount; i++ {
		articles += getArticle(files[i])
	}

	doc.Find("#content").AppendHtml(articles)

	docHtmlStr, err := doc.Selection.Html()
	check(err)

	filename := fmt.Sprintf("build/combined/[%d-%d] "+DOMAIN+".html", theRange.iMin, theRange.iMax)

	osFile, err := os.Create(filename)

	check(err)

	fmt.Println("Generated: " + filename)

	osFile.WriteString(docHtmlStr)

	osFile.Close()
}

func getArticle(htmlFile HtmlFile) string {

	var article string
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(htmlFile.Content)))

	if err != nil {
		panic(err)
	}

	article, err = doc.Find("div#content").Html()

	if err != nil {
		panic(err)
	}

	return article
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
