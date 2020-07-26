package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/thejini3/blog-to-pdf/sitemap"
	hel "github.com/thejini3/go-helper"
)

func getHTMLFiles() []htmlFileStruct {

	var htmlFiles []htmlFileStruct
	var urls = getUrls()

	// p("getUrls count: " + strconv.Itoa(len(getUrls())))

	for i, urlStr := range urls {

		// fileNo := strconv.Itoa(i + 1)

		if cfg.LimitUrlsNo > 0 && (i+1) > cfg.LimitUrlsNo {
			break
		}

		localHTMLFilePath := originalHTMLDir + "/" + hel.AZ_AND_NUMBER_ONLY(urlStr) + ".html"

		if cfg.ForceFetchHTML || !hel.FileExists(localHTMLFilePath) {

			osFile, err := os.Create(localHTMLFilePath)

			if err != nil {
				panic(err)
			}

			urlContent := hel.GetURLContent(urlStr, cfg.BrowserUserAgent)
			osFile.WriteString(string(urlContent))

			hel.P(fmt.Sprintf("%v: Downloaded Origin Html: %v", i+1, localHTMLFilePath))

			osFile.Close()
		}

		htmlFiles = append(htmlFiles, htmlFileStruct{
			LocalPath: localHTMLFilePath,
			RemoteURL: urlStr,
		})

	}
	// pp("htmlFiles count: " + strconv.Itoa(len(htmlFiles)))
	//hel.P("Run again, if app quits here!")
	return htmlFiles
}

func getUrls() []string {

	if !cfg.ForceUrlsFetch && hel.FileExists(cfg.URLFile) == true {

		urls := hel.StrToArr(string(hel.GetFileBytes(cfg.URLFile)), "\n")

		if cfg.LimitUrlsNo > 0 {
			return hel.LimitStrArr(urls, cfg.LimitUrlsNo)
		}

		return urls
	}

	var urlStr = getUrlsFromSiteMap()

	if len(urlStr) == 0 {
		hel.PS("SiteMap url isn't valid probably!")
		hel.PM("So, try following wget command go grab all url of the site")
		hel.PM(`$ wget --spider -r https://` + cfg.Domain + ` 2>&1 | grep '^--' | awk '{ print $3 }' | grep -v '\.\(css\|js\|png\|gif\|jpg\|JPG\)$' > /tmp/urls.txt`)
		hel.PE("Then copy full content of 'wget.urls.txt' file to urls.txt")
	}

	f, err := os.Create(cfg.URLFile)

	if err != nil {
		fmt.Println("Error os.Create: "+err.Error(), cfg.URLFile)
	}

	defer f.Close()

	f.WriteString(urlStr)

	urls := hel.StrToArr(urlStr, "\n")

	if cfg.LimitUrlsNo > 0 {
		return hel.LimitStrArr(urls, cfg.LimitUrlsNo)
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
