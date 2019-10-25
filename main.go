package main

import (
	"fmt"
)

var protoCol = "https://"

var buildDir string
var pdfDir string
var htmlDir string
var urlsTxtPath string

var ArticlePerPDF = 10
var forceFetchHtml bool
var isTestRun bool
var generatePdf bool
var pdfPageSize string
var DOMAIN string
var isHttps bool
var SiteURL string
var forceUrlsFetch bool
var SiteMapURL string

func main() {

	bootFlag()
	bootPaths()

	p("booted")

	buildAllHTMLS()

}

func check(err error) {
	if err != nil {
		fmt.Println("err: ", err.Error())
		panic(err)
	}
}
