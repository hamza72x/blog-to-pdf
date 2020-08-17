package main

import (
	"fmt"
	"os"
	"sort"
	"sync"

	"github.com/thejini3/blog-to-pdf/sitemap"
	hel "github.com/thejini3/go-helper"
)

func getHTMLFiles() []xHTMLFile {

	var htmlFiles []xHTMLFile
	var urls = getUrls()

	var wg sync.WaitGroup
	var c = make(chan int, thread)

	for i, urlStr := range urls {

		if cfg.LimitUrlsNo > 0 && (i+1) > cfg.LimitUrlsNo {
			break
		}

		wg.Add(1)

		localHTMLFilePath := originalHTMLDir + "/" + hel.StrFilterToAlphabetsAndNumbersMust(urlStr) + ".html"

		if cfg.ForceFetchHTML || !hel.FileExists(localHTMLFilePath) {

			go func(localHTMLFilePath, urlStr string, i int) {
				c <- i
				download(localHTMLFilePath, urlStr, i)
				wg.Done()
				<-c
			}(localHTMLFilePath, urlStr, i)

		} else {
			wg.Done()
		}

		htmlFiles = append(htmlFiles, xHTMLFile{
			LocalPath: localHTMLFilePath,
			RemoteURL: urlStr,
		})

	}

	wg.Wait()

	return htmlFiles
}
func download(localHTMLFilePath, urlStr string, i int) {
	osFile, err := os.Create(localHTMLFilePath)

	if err != nil {
		panic(err)
	}

	urlContent := hel.URLContentMust(urlStr, cfg.BrowserUserAgent)
	osFile.WriteString(string(urlContent))

	hel.Pl(fmt.Sprintf("%v: Downloaded Origin Html: %v", i+1, localHTMLFilePath))

	osFile.Close()
}
func getUrls() []string {

	if !cfg.ForceUrlsFetch && hel.FileExists(cfg.URLFile) == true {

		urls := hel.StrToArr(string(hel.FileBytesMust(cfg.URLFile)), "\n")

		if cfg.LimitUrlsNo > 0 {
			return hel.StrArrLimit(urls, cfg.LimitUrlsNo)
		}

		return urls
	}

	var urlStr = getUrlsFromSiteMap()

	if len(urlStr) == 0 {
		panic("SiteMap url isn't valid probably, use $ sitemap-generator (npm package)")
	}

	if err := hel.StrToFile(cfg.URLFile, urlStr); err != nil {
		panic("Error writing file to: " + cfg.URLFile)
	}

	urls := hel.StrToArr(urlStr, "\n")

	if cfg.LimitUrlsNo > 0 {
		return hel.StrArrLimit(urls, cfg.LimitUrlsNo)
	}

	return urls
}

func getUrlsFromSiteMap() string {

	var allUrls = ""

	for _, url := range cfg.SiteMapsURL {

		var smap sitemap.Sitemap
		var err error

		if cfg.GetSiteMapByWpJSON {

			hel.Pl("Getting urls by wp-json")

			smap = sitemap.GetByWPJSON(url, cfg.BrowserUserAgent)

		} else {

			hel.Pl("Getting urls by standard sitemap")

			smap, err = sitemap.Get(url, nil)

			if err != nil {
				fmt.Println("Site map get error: " + err.Error())
			}
		}

		for _, url := range getSortedSiteMapURL(smap.URLS) {
			if ignoreURL(url.Loc) {
				continue
			}
			allUrls += url.Loc + "\n"
		}
	}

	return allUrls
}

func ignoreURL(urlStr string) bool {
	u1 := "https://" + cfg.Domain
	u2 := "http://" + cfg.Domain
	c1 := (urlStr == u1) || (urlStr == u1+"/") || (urlStr == u1+"/about") || (urlStr == u1+"/contact")
	c2 := (urlStr == u2) || (urlStr == u2+"/") || (urlStr == u2+"/about") || (urlStr == u2+"/contact")
	return c1 || c2
}

func getSortedSiteMapURL(urls []sitemap.URL) []sitemap.URL {
	if cfg.PostOrder == "asc" {
		sort.Slice(urls, func(i, j int) bool {
			return urls[i].GetTime().Before(urls[j].GetTime())
		})
	} else {
		sort.Slice(urls, func(i, j int) bool {
			return urls[i].GetTime().After(urls[j].GetTime())
		})
	}
	return urls
}
