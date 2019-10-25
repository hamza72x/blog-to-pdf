package main

import (
	"os"
)

const UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1"

var SiteBasedTags = map[string][]string{
	"alorpothe.wordpress.com": {"div#secondary", "div.menu-search", "nav#nav-single",
		"footer#colophon", "div.widget", "div#fb-root", "div#actionbar",
		"form", "h3#reply-title", "div#jp-post-flair", "div#header-img", ".cs-rating .pd-rating",
		"div.wpcnt", "h3#entry-format", ".rating-star-icon",
	},
}

var protocol = "https://"

var buildDir string
var htmlDir string
var siteMapFilePath string

var ArticlePerPDF = 10
var forceFetchHtml bool
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
		panic(err)
	}
}
