package main

import (
	"fmt"
	"os"
	"sort"
	"github.com/thejini3/blog-to-pdf/sitemap"
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

	// p("getUrls count: " + strconv.Itoa(len(getUrls())))

	for i, urlStr := range urls {

		// fileNo := strconv.Itoa(i + 1)

		if cfg.LimitUrlsNo > 0 && (i+1) > cfg.LimitUrlsNo {
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
			LocalPath: localHtmlFilePath,
			Content:   removeTags(getFileBytes(localHtmlFilePath)),
			URL:       urlStr,
		})

	}
	// pp("htmlFiles count: " + strconv.Itoa(len(htmlFiles)))
	p("Run again, if app quits here!")
	return htmlFiles
}

func getUrls() []string {

	if !cfg.ForceUrlsFetch && fileExists(urlsTxtFilePath) == true {

		urls := strToArr(string(getFileBytes(urlsTxtFilePath)), "\n")

		if cfg.LimitUrlsNo > 0 {
			return limitStrArr(urls, cfg.LimitUrlsNo)
		}

		return urls
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

	urls := strToArr(urlStr, "\n")

	if cfg.LimitUrlsNo > 0 {
		return limitStrArr(urls, cfg.LimitUrlsNo)
	}

	return urls
}

func getUrlsFromSiteMap() string {

	var allUrls = ""

	smap, err := sitemap.Get(cfg.SiteMapURL, nil)

	if err != nil {
		fmt.Println("Site map get error: " + err.Error())
	}

	for _, url := range getSortedSiteMapURL(smap.URLS) {
		if ignoreURL(url.Loc) {
			continue
		}
		allUrls += url.Loc + "\n"
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
