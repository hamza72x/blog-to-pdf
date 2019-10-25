package main

import (
	"os"
	"fmt"
)

const UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1"


var protoCol = "https://"

var buildDir string
var pdfDir string
var htmlDir string
var siteMapFilePath string

var ArticlePerPDF = 10
var forceFetchHtml bool
var pdfPageSize string
var DOMAIN string
var isHttps bool
var SiteURL string
var forceSiteMapFetch bool
var SiteMapURL string

func main() {

	bootFlag()
	bootPaths()
	generateAllPdf()
	os.Exit(0)
}

func check(err error) {
	if err != nil {
		fmt.Println("err: ", err.Error())
		panic(err)
	}
}
