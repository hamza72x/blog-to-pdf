package main

import hel "github.com/thejini3/go-helper"

type iniStruct struct {
	Domain                         string   `ini:"domain"`
	SiteMapsURL                    []string `ini:"sitemaps_url"`
	GetSiteMapByWpJSON             bool     `ini:"get_sitemap_by_wp_json"`
	URLFile                        string   `ini:"url_file"`
	BrowserUserAgent               string   `ini:"browser_user_agent"`
	PdfFileName                    string   `ini:"pdf_file_name"`
	PostOrder                      string   `ini:"post_order"`
	ArticlePerPDF                  int      `ini:"article_per_pdf"`
	ArticleParentElement           string   `ini:"article_parent_element"`
	ArticleTitleClass              string   `ini:"article_title_class"`
	AppendURLInTitle               bool     `ini:"append_article_url_in_title"`
	AppendAutoArticleNumberInTitle bool     `ini:"append_auto_article_number_in_title"`
	ElementsToRemove               []string `ini:"elements_to_remove"`
	StringReplacesFile             string   `ini:"string_replaces_file"`
	PatternReplacesFile            string   `ini:"pattern_replaces_file"`
	ForceFetchHTML                 bool     `ini:"force_html_fetch"`
	ForceUrlsFetch                 bool     `ini:"force_urls_fetch"`
	LimitUrlsNo                    int      `ini:"limit_urls"`
	GeneratePDF                    bool     `ini:"generate_pdf"`
	SkipPDFCreationIfExistsAlready bool     `ini:"skip_pdf_creation_if_exists_already"`
	PdfPageSize                    string   `ini:"pdf_size"`
	PdfOrientation                 string   `ini:"pdf_orientation"`
	CustomCSSFile                  string   `ini:"custom_css_file"`
	PdfOutputDirPath               string   `ini:"pdf_output_dir_path"`
	PdfMarginTop                   int      `ini:"pdf_margin_top"`
	PdfMarginLeft                  int      `ini:"pdf_margin_left"`
	PdfMarginRight                 int      `ini:"pdf_margin_right"`
	PdfMarginBottom                int      `ini:"pdf_margin_bottom"`
}

type rangeStruct struct {
	Min int // 1 based, i.e not 0 based -_-
	Max int
}

type htmlFileStruct struct {
	LocalPath string
	RemoteURL string
}

type singleOutFileStruct struct {
	HTMLFiles []htmlFileStruct
	TheRange  rangeStruct
	FileNo    int
}

func (hf *htmlFileStruct) HtmlBytes() []byte {
	return []byte(hel.GetFileBytes(hf.LocalPath))
}

type theReplace struct {
	Serial int               `json:"serial"`
	Data   map[string]string `json:"data"`
}

func (t *theReplace) FindStr() string {
	var str = ""
	for i := range t.Data {
		str = i
	}
	return str
}

func (t *theReplace) ReplaceStr() string {
	var str = ""
	for i := range t.Data {
		str = t.Data[i]
	}
	return str
}

//
//func (sof *singleOutFileStruct) combinedHtmlStr() string {
//
//	firstHtmlFile := sof.OutArticles[0].ContentHtml
//
//	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(firstHtmlFile.Bytes()))
//	hel.PErr("goquery.NewDocumentFromReader", err)
//
//	head := doc.Find("head")
//
//	head.AppendHtml("<style>" + ConstDefaultCss + "</style>")
//
//	if hel.FileExists(cfg.CustomCSSFile) {
//		head.AppendHtml(`<style>` + hel.GetFileStr(cfg.CustomCSSFile) + `</style>`)
//	}
//
//	// combined htmls - ArticlePerPDF
//	for i := 1; i < len(combinedHtmlFiles); i++ {
//		articleParent := doc.Find(cfg.ArticleParentElement)
//		articleParent.AppendHtml(`<br/><hr/><hr/><br/>`)
//		articleParent.AppendHtml(getContentHtml(combinedHtmlFiles[i]))
//	}
//
//	// set [i] in title
//	doc.Find(cfg.ArticleTitleClass).Each(func(i int, s *goquery.Selection) {
//		s.PrependHtml("[" + strconv.Itoa(theRange.Min+i) + "] ")
//		s.AddClass("text-center")
//		if cfg.AppendURLInTitle {
//			s.AppendHtml(
//				fmt.Sprintf("<br/><a class=\"article-origin-link\" style=\"font-size: 12px;\" href=\"%s\">%s</a>", combinedHtmlFiles[i].RemoteURL, combinedHtmlFiles[i].RemoteURL),
//			)
//		}
//	})
//
//	docHtmlStr, err := doc.Selection.Html()
//	hel.PErr("doc.Selection.Html", err)
//}
