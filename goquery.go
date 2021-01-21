package main

import (
	"bytes"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	hel "github.com/hamza02x/go-helper"
	"github.com/jaytaylor/html2text"
)

// func getContentTxt(file xHTMLFile) string {
// 	text, err := html2text.FromString(string(file.Bytes()), html2text.Options{PrettyTables: true})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return text
// }

func getContentHTML(htmlFile xHTMLFile) string {

	var content string
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlFile.fileBytes()))
	hel.PlP("Error in getting content [NewDocumentFromReader]", err)

	var articleParent = doc.Find(cfg.ArticleParentElement)
	// articleParent.Children().First().SetAttr("style", "page-break-before: always;")

	content, err = articleParent.Html()

	if err != nil {
		panic(err)
	}

	var re = regexp.MustCompile(`\<script .*\>([^@]*)\<\/script\>`)

	return additionalFilter(re.ReplaceAllString(content, ""))
}

func getTitleTxt(htmlFile xHTMLFile) string {

	text, err := html2text.FromString(getTitleHTML(htmlFile), html2text.Options{PrettyTables: true})
	if err != nil {
		panic(err)
	}

	return text
}

func getTitleHTML(htmlFile xHTMLFile) string {

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlFile.fileBytes()))
	hel.PlP("Error in getting title [NewDocumentFromReader]", err)

	htmlStr, err := doc.Find(cfg.ArticleTitleClass).Html()
	hel.PlP("Error in getting doc.Find(cfg.ArticleTitleClass).Html()", err)

	return htmlStr
}
