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

	for _, tag := range cfg.ElementsToRemove {
		doc.Find(tag).Remove()
	}

	htmlStr, err := doc.Html()

	if err != nil {
		panic(err)
	}

	for _, strReplace := range strReplaces {
		htmlStr = strings.ReplaceAll(htmlStr, strReplace.ReplaceKey, strReplace.ReplaceVal)
	}

	return htmlStr

}
