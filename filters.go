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

	//for _, element := range DefaultElementsToBeRemoved {
	//	doc.Find(element).Remove()
	//}

	for _, tag := range elementsToRemove {
		doc.Find(tag).Remove()
	}

	htmlStr, err := doc.Html()

	if err != nil {
		panic(err)
	}

	//for k, v := range DefaultStringsToBeReplaced {
	//	htmlStr = strings.ReplaceAll(htmlStr, k, v)
	//}

	//specialReplaces, ok2 := SpecialStringsReplaceAll[siteDomain]
	//
	//if ok2 {
	//	for _, replaceMaps := range specialReplaces {
	//		for k, v := range replaceMaps {
	//			htmlStr = strings.ReplaceAll(htmlStr, k, v)
	//		}
	//	}
	//}

	return htmlStr

}
