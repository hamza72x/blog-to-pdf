package main

import (
	"bytes"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/jaytaylor/html2text"
)

func getContentHTML(htmlFile xHTMLFile) (string, error) {

	var content string
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlFile.fileBytes()))

	if err != nil {
		return "", err
	}

	var articleParent = doc.Find(cfg.ArticleParentElement)

	content, err = articleParent.Html()

	if err != nil {
		return "", err
	}

	var re = regexp.MustCompile(`\<script .*\>([^@]*)\<\/script\>`)

	return additionalFilter(re.ReplaceAllString(content, "")), nil
}

func getTitleTxt(htmlFile xHTMLFile) (string, error) {

	title, err := getTitleHTML(htmlFile)

	if err != nil {
		return "", err
	}

	text, err := html2text.FromString(title, html2text.Options{PrettyTables: true})

	if err != nil {
		return "", err
	}

	return text, nil
}

func getTitleHTML(htmlFile xHTMLFile) (string, error) {

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlFile.fileBytes()))

	if err != nil {
		return "", err
	}

	htmlStr, err := doc.Find(cfg.ArticleTitleClass).Html()

	if err != nil {
		return "", err
	}

	return htmlStr, nil
}
