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

		if isTestRun && i == 11 {
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

	if !forceUrlsFetch {
		if fileExists(urlsTxtPath) {
			return strings.Split(
				strings.ReplaceAll(string(getFileContents(urlsTxtPath)), " ", ""),
				"\n",
			)
		}
	}

	f, err := os.Create(urlsTxtPath)

	if err != nil {
		fmt.Println("Error os.Create: "+err.Error(), urlsTxtPath)
	}

	if !ContainsStr(siteMapNotApplicables, DOMAIN) {
		f.WriteString(getUrlsFromSiteMap())
	} else {
		// have to get from feed I guess
		urlContent := getURLContent(SiteURL + "/feed")
		p(string(urlContent))
		panic("NO SITE MAP DUDE!")
	}

	f.Close()

	return urls
}

func getUrlsFromSiteMap() (allUrls string) {

	smap, err := sitemap.Get(SiteMapURL, nil)

	if err != nil {
		fmt.Println("Site map get error: " + err.Error())
	}

	if err != nil {
		fmt.Println("Error os.Create: "+err.Error(), urlsTxtPath)
	}

	var iCount = len(smap.URLS) - 1

	for i, url := range getSortedSiteMapURL(smap.URLS) {

		if ignoreURL(url.Loc) {
			continue
		}

		if isTestRun && i == 11 {
			break
		}

		if iCount == i {
			allUrls += url.Loc
		} else {
			allUrls += url.Loc + "\n"
		}

	}
	return
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
