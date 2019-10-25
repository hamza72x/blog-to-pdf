package main

import (
	"github.com/yterajima/go-sitemap"
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"net/http"
	"time"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

type HtmlFile struct {
	Name      string
	LocalPath string
	Content   string
	URL       string
}

func getHtmlFiles(forceFetchHTMLS bool, forceFetchSiteMap bool) []HtmlFile {

	var htmlFiles []HtmlFile

	for _, urlStr := range getUrls(false) {

		path := "htmls/" + removeSpecialChars(urlStr) + ".html"

		if forceFetchHTMLS || !fileExists(path) {

			osFile, err := os.Create(path)

			if err != nil {
				panic(err)
			}

			osFile.WriteString(
				removeTags(
					getURLContent(urlStr),
				),
			)

			fmt.Println("Did Write: " + path)

			osFile.Close()
		}

		htmlFiles = append(htmlFiles, HtmlFile{
			Name:      removeSpecialChars(urlStr),
			LocalPath: path,
			Content:   getFileContentAsString(path),
			URL:       urlStr,
		})

	}

	return htmlFiles
}

func getUrls(forceFetchSiteMap bool) []string {

	var urls []string
	localSiteMap := getLocalSiteMapUrlsFilePath()

	if !forceFetchSiteMap {
		if fileExists(localSiteMap) {
			return strings.Split(getFileContentAsString(strings.ReplaceAll(localSiteMap, " ", "")), "\n")
		}
	}

	smap, err := sitemap.Get(SiteMapURL, nil)

	if err != nil {
		fmt.Println("Site map get error: " + err.Error())
	}

	f, err := os.Create(localSiteMap)

	if err != nil {
		fmt.Println("Error os.Create: "+err.Error(), localSiteMap)
	}

	var iCount = len(smap.URL) - 1

	for i, url := range smap.URL {
		if ignoreURL(url.Loc) {
			continue
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
	return (urlStr == SITE) || (urlStr == SITE+"/")
}

func getFileContentAsString(filePath string) string {

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error reading file: " + filePath)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("Error ioutil.ReadAll: " + filePath)
	}

	return string(b)
}

func fileExists(filename string) bool {

	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getLocalSiteMapUrlsFilePath() string {
	return "urls/" + removeSpecialChars(SiteMapURL) + ".txt"
}

func removeSpecialChars(str string) string {
	str = strings.ReplaceAll(str, "/", "-")
	str = strings.ReplaceAll(str, ":", "-")
	return str
}

func getURLContent(urlStr string) []byte {

	// fmt.Printf("HTML code of %s ...\n", urlStr)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", urlStr, nil)

	if err != nil {
		panic(err)
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1")

	// Make request
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	htmlBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return htmlBytes
}

func removeTags(htmlBytes []byte) string {

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlBytes))

	if err != nil {
		panic(err)
	}

	tags, ok := SiteBasedTags[DOMAIN]
	if ok {
		for _, tag := range tags {
			doc.Find(tag).Each(func(i int, s *goquery.Selection) {
				s.Remove()
			})
		}
	}



	//doc.Find("head").Each(func(i int, s *goquery.Selection) {
	//	s.Append("<style> body p { font-family: \"Kohinoor Bangla\", serif !important; font-size: 20px !important; } </style>")
	//})

	htmlStr, err := doc.Html()

	if err != nil {
		panic(err)
	}

	return htmlStr

}
