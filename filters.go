package main

import (
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

func removeTags(htmlBytes []byte) string {

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlBytes))

	if err != nil {
		panic(err)
	}

	tags, ok := SiteBasedTags[DOMAIN]

	if ok {
		for _, tag := range tags {
			doc.Find(tag).Remove()
		}
	}

	htmlStr, err := doc.Html()

	if err != nil {
		panic(err)
	}

	return htmlStr

}
