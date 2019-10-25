package main

import (
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"strings"
)

func removeTags(htmlBytes []byte) string {

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlBytes))

	if err != nil {
		panic(err)
	}

	for _, element := range DefaultElementsToBeRemoved {
		doc.Find(element).Remove()
	}

	tags, ok := DivsToBeRemoved[DOMAIN]

	if ok {
		for _, tag := range tags {
			doc.Find(tag).Remove()
		}
	}

	htmlStr, err := doc.Html()

	if err != nil {
		panic(err)
	}

	for k, v := range DefaultStringsToBeReplaced {
		htmlStr = strings.ReplaceAll(htmlStr, k, v)
	}

	specialReplaces, ok2 := SpecialStringsReplaceAll[DOMAIN]

	if ok2 {
		for _, replaceMaps := range specialReplaces {
			for k, v := range replaceMaps {
				htmlStr = strings.ReplaceAll(htmlStr, k, v)
			}
		}
	}

	return htmlStr

}
