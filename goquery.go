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

	content, err = doc.Find(articleParentElement).Html()

	if err != nil {
		panic(err)
	}

	return content
}
