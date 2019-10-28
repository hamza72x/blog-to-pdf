package main

import (
	"fmt"
	"os"
	"sort"
	"gitlab.com/thejini3/blog-to-pdf/sitemap"
	"strconv"
)

type HtmlFile struct {
	// Name      string
	LocalPath string
	Content   string
	URL       string
}

func getHtmlFiles() []HtmlFile {

	var htmlFiles []HtmlFile
	var urls = getUrls()

	for i, urlStr := range urls {

		if cfg.limitUrls > 0 && i == cfg.limitUrls-1 {
			break
		}

		localHtmlFilePath := originalHtmlDir + "/" + strconv.Itoa(i+1) + ".html"

		if cfg.ForceFetchHtml || !fileExists(localHtmlFilePath) {

			osFile, err := os.Create(localHtmlFilePath)

			if err != nil {
				panic(err)
			}

			urlContent := getURLContent(urlStr)
			osFile.WriteString(string(urlContent))

			p(fmt.Sprintf("%v: Downloaded Origin Html: %v\n", i+1, localHtmlFilePath))

			osFile.Close()
		}

		htmlFiles = append(htmlFiles, HtmlFile{
			// Name:      getHtmlLocalFileNameFromUrl(urlStr),
			LocalPath: localHtmlFilePath,
			Content:   removeTags(getFileContents(localHtmlFilePath)),
			URL:       urlStr,
		})

	}
	p("Run again if app quits here!")
	return htmlFiles
}

func getUrls() []string {

	if !cfg.ForceUrlsFetch && fileExists(urlsTxtFilePath) == true {
		return strToArr(string(getFileContents(urlsTxtFilePath)), "\n")
	}

	var urlStr = getUrlsFromSiteMap()

	if len(urlStr) == 0 {
		ps("SiteMap url isn't valid probably!")
		pm("So, try following wget command go grab all url of the site")
		pm(`$ wget --spider -r ` + SiteURL + ` 2>&1 | grep '^--' | awk '{ print $3 }' | grep -v '\.\(css\|js\|png\|gif\|jpg\|JPG\)$' > /tmp/urls.txt`)
		pe("Then copy full content of 'wget.urls.txt' file to urls.txt")
	}

	f, err := os.Create(urlsTxtFilePath)

	if err != nil {
		fmt.Println("Error os.Create: "+err.Error(), urlsTxtFilePath)
	}

	defer f.Close()

	f.WriteString(urlStr)

	return strToArr(urlStr, "\n")
}

func getUrlsFromSiteMap() string {
	var allUrls = ""

	smap, err := sitemap.Get(cfg.SiteMapURL, nil)

	if err != nil {
		fmt.Println("Site map get error: " + err.Error())
	}

	var iCount = len(smap.URLS) - 1

	for i, url := range getSortedSiteMapURL(smap.URLS) {

		if ignoreURL(url.Loc) {
			continue
		}

		if cfg.limitUrls > 0 && i == cfg.limitUrls-1 {
			break
		}

		if iCount == i {
			allUrls += url.Loc
		} else {
			allUrls += url.Loc + "\n"
		}

	}
	return allUrls
}

func ignoreURL(urlStr string) bool {
	return (urlStr == SiteURL) || (urlStr == SiteURL+"/") || (urlStr == SiteURL+"/about") || (urlStr == SiteURL+"/contact")
}

func getSortedSiteMapURL(urls []sitemap.URL) []sitemap.URL {

	sort.Slice(urls, func(i, j int) bool {
		return urls[i].GetTime().Before(urls[j].GetTime())
	})
	return urls
}
