package main

import (
	"bytes"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/jaytaylor/html2text"
	hel "github.com/thejini3/go-helper"
)

// func getContentTxt(file htmlFileStruct) string {
// 	text, err := html2text.FromString(string(file.Bytes()), html2text.Options{PrettyTables: true})
// 	if err != nil {
// 		panic(err)
// 	}
// 	return text
// }

func getContentHTML(htmlFile htmlFileStruct) string {

	var content string
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlFile.HtmlBytes()))
	hel.PErr("Error in getting content [NewDocumentFromReader]", err)

	var articleParent = doc.Find(cfg.ArticleParentElement)
	// articleParent.Children().First().SetAttr("style", "page-break-before: always;")

	content, err = articleParent.Html()

	if err != nil {
		panic(err)
	}

	var re = regexp.MustCompile(`\<script .*\>([^@]*)\<\/script\>`)

	return additionalFilter(re.ReplaceAllString(content, ""))
}

func getTitleTxt(htmlFile htmlFileStruct) string {

	text, err := html2text.FromString(getTitleHTML(htmlFile), html2text.Options{PrettyTables: true})
	if err != nil {
		panic(err)
	}

	return text
}

func getTitleHTML(htmlFile htmlFileStruct) string {

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlFile.HtmlBytes()))
	hel.PErr("Error in getting title [NewDocumentFromReader]", err)

	htmlStr, err := doc.Find(cfg.ArticleTitleClass).Html()

	return htmlStr
}
