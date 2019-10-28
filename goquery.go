package main

import (
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

func getContent(htmlFile HtmlFile) string {

	var content string
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte(htmlFile.Content)))

	if err != nil {
		panic(err)
	}

	var articleParent = doc.Find(cfg.ArticleParentElement)
	articleParent.Children().First().AddClass(ConstPageBreakClass)

	content, err = articleParent.Html()

	if err != nil {
		panic(err)
	}

	return content
}
