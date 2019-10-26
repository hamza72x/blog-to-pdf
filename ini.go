package main

import (
	"os"
	"strings"
	"fmt"
)

var siteDomain string
var articleParentElement string
var SiteURL string
var SiteMapURL string
var pdfPageSize string
var isTestRun bool
var forceFetchHtml bool
var forceUrlsFetch bool
var generatePDF bool
var pdfMarginTop int
var pdfMarginLeft int
var pdfMarginRight int
var pdfMarginBottom int
var elementsToRemove []string

func generateIniFile() {

	f, err := os.Create("configs.any_blog.com.ini")

	if err != nil {
		pp("Error creating ini file: " + err.Error())
	}
	defer f.Close()

	f2, err2 := os.Create("string_replaces.any_blog.com.json")

	if err2 != nil {
		pp("Error creating ini file: " + err2.Error())
	}

	defer f2.Close()

	f2.WriteString(ConstReplaces)
	f.WriteString(ConstSampleINI)
}

func bootIni() {

	siteDomain = iniGetStr("domain")
	articleParentElement = iniGetStr("article_parent_element")
	pdfPageSize = iniGetStr("pdf_size")
	SiteMapURL = iniGetStr("sitemap_url")
	isTestRun = iniGetBool("is_test_run")
	forceFetchHtml = iniGetBool("force_html_fetch")
	generatePDF = iniGetBool("generate_pdf")
	forceUrlsFetch = iniGetBool("force_urls_fetch")
	SiteURL = strings.ReplaceAll(iniGetStr("protocol")+siteDomain, "/", "")
	pdfMarginTop = iniGetInt("pdf_margin_top")
	pdfMarginLeft = iniGetInt("pdf_margin_left")
	pdfMarginRight = iniGetInt("pdf_margin_right")
	pdfMarginBottom = iniGetInt("pdf_margin_bottom")
	elementsToRemove = strings.Split(strings.TrimSpace(iniGetStr("elements_to_remove")), ",")

	fmt.Println("elementsToRemove", elementsToRemove)

	p("Initialized all ini config!")
}

func iniGetStr(iniKey string) string {
	return cfg.Section("").Key(iniKey).String()
}
func iniGetInt(iniKey string) int {
	val, err := cfg.Section("").Key(iniKey).Int()
	check(err, iniKey)
	return val
}
func iniGetBool(iniKey string) bool {
	val, err := cfg.Section("").Key(iniKey).Bool()
	check(err, iniKey)
	return val
}

func check(err error, iniKey string) {
	if err != nil {
		pp("Error ini data: " + iniKey)
	}
}
