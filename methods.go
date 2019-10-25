package main

import (
	"fmt"
	"os"
	"strings"
	"sort"
	"gitlab.com/thejini3/blog-to-pdf/sitemap"
)

type HtmlFile struct {
	Name      string
	LocalPath string
	Content   string
	URL       string
}

func getHtmlFiles() []HtmlFile {

	var htmlFiles []HtmlFile

	for i, urlStr := range getUrls() {

		if isTestRun, i == 11 {
			break
		}

		path := htmlDir + "/" + removeSpecialChars(urlStr) + ".html"

		if forceFetchHtml || !fileExists(path) {

			osFile, err := os.Create(path)

			if err != nil {
				panic(err)
			}

			osFile.Write(getURLContent(urlStr))

			p(fmt.Sprintf("%v: Did Write: %v\n", i+1, path))

			osFile.Close()
		}

		htmlFiles = append(htmlFiles, HtmlFile{
			Name:      removeSpecialChars(urlStr),
			LocalPath: path,
			Content:   removeTags(getFileContents(path)),
			URL:       urlStr,
		})

	}

	return htmlFiles
}

func getUrls() []string {

	var urls []string
	// localSiteMap := getLocalSiteMapUrlsFilePath()

	if !forceSiteMapFetch {
		if fileExists(siteMapFilePath) {
			return strings.Split(
				strings.ReplaceAll(string(getFileContents(siteMapFilePath)), " ", ""),
				"\n",
			)
		}
	}

	smap, err := sitemap.Get(SiteMapURL, nil)

	if err != nil {
		fmt.Println("Site map get error: " + err.Error())
	}

	f, err := os.Create(siteMapFilePath)

	if err != nil {
		fmt.Println("Error os.Create: "+err.Error(), siteMapFilePath)
	}

	var iCount = len(smap.URLS) - 1

	for i, url := range getSortedSiteMapURL(smap.URLS) {

		if ignoreURL(url.Loc) {
			continue
		}

		if isTestRun, i == 11 {
			break
		}

		if iCount == i {
			f.WriteString(url.Loc)
		} else {
			f.WriteString(url.Loc + "\n")
		}
		urls = append(urls, url.Loc)
	}

	f.Close()

	return urls
}

func ignoreURL(urlStr string) bool {
	return (urlStr == SiteURL) || (urlStr == SiteURL+"/")
}

func getSortedSiteMapURL(urls []sitemap.URL) []sitemap.URL {

	sort.Slice(urls, func(i, j int) bool {
		return urls[i].GetTime().Before(urls[j].GetTime())
	})
	return urls
}
