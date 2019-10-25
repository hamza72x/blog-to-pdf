package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func bootFlag() {

	flag.BoolVar(&forceFetchHtml, "force-html-fetch", false,
		"Re-fetch htmls from server if it's not already fetched in local directory",
	)
	flag.BoolVar(&isTestRun, "is-test-run", false,
		"if yes, then it will fetch only 10 url to test!",
	)
	flag.BoolVar(&forceUrlsFetch, "force-urls-fetch", false,
		"Re-fetch htmls from server if it's not already fetched in local directory",
	)
	flag.StringVar(&DOMAIN, "domain", "-",
		"(Required) Domain of the site, Ex: alorpothe.wordpress.com",
	)
	flag.BoolVar(&isHttps, "https", true,
		"https or not",
	)
	flag.BoolVar(&generatePdf, "generate-pdf", true,
		"Generate pdf or not, then just html will be created!",
	)
	flag.IntVar(&ArticlePerPDF, "article-per-pdf", 10,
		"The number of articles per pdf",
	)
	flag.IntVar(&pdfMargin, "pdf-margin", 3,
		"Margin around the contents of PDF",
	)
	flag.StringVar(&pdfPageSize, "pdf-size", "A7",
		"The size of output PDF",
	)
	flag.Parse()

	if checkDomain(DOMAIN) != nil {
		p("Wrong, domain name: " + checkDomain(DOMAIN).Error())
		flag.Usage()
		os.Exit(0)
	}

	DOMAIN = strings.ReplaceAll(DOMAIN, "/", "")

	if !isHttps {
		protoCol = "http://"
	}

	SiteURL = protoCol + DOMAIN
	SiteMapURL = SiteURL + "/sitemap.xml"

	p("Current Configs")
	fmt.Printf("-article-per-pdf: %v\n", ArticlePerPDF)
	fmt.Println("-domain: " + DOMAIN)
	fmt.Printf("-force-html-fetch: %v\n", forceFetchHtml)
	fmt.Printf("-force-urls-fetch: %v\n", forceUrlsFetch)
	fmt.Printf("-generate-pdf: %v\n", generatePdf)
	fmt.Printf("-https: %v\n", isHttps)
	fmt.Printf("-pdf-margin: %v\n", pdfMargin)

	fmt.Printf("SiteURL: %v\n", SiteURL)

	if !ContainsStr(siteMapNotApplicables, DOMAIN) {
		fmt.Printf("SiteMapURL: %v\n", SiteMapURL)
	}
}
