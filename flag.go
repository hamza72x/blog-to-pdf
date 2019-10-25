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
	flag.BoolVar(&forceSiteMapFetch, "force-sitemap-fetch", false,
		"Re-fetch htmls from server if it's not already fetched in local directory",
	)
	flag.StringVar(&DOMAIN, "domain", "",
		"(Required) Domain of the site, Ex: alorpothe.wordpress.com",
	)
	flag.BoolVar(&isHttps, "https", true,
		"https or not",
	)
	flag.IntVar(&ArticlePerPDF, "article-per-pdf", 10,
		"The number of articles per pdf",
	)

	flag.Parse()

	DOMAIN = strings.ReplaceAll(DOMAIN, "/", "")

	if checkDomain(DOMAIN) != nil {
		fmt.Println("Wrong, domain name: " + checkDomain(DOMAIN).Error())
		flag.Usage()
		os.Exit(0)
	}

	if !isHttps {
		protocol = "http://"
	}

	SiteURL = protocol + DOMAIN
	SiteMapURL = SiteURL + "/sitemap.xml"

	fmt.Printf("Current Configs - \n\n")
	fmt.Printf("article-per-pdf: %v\n", ArticlePerPDF)
	fmt.Println("DOMAIN: " + DOMAIN)
	fmt.Printf("force-html-fetch: %v\n", forceFetchHtml)
	fmt.Printf("force-sitemap-fetch: %v\n", forceSiteMapFetch)
	fmt.Printf("isHttps: %v\n", isHttps)
	fmt.Printf("SiteURL: %v\n", SiteURL)
	fmt.Printf("SiteMapURL: %v\n", SiteMapURL)

}
