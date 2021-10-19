package main

import hel "github.com/hamza72x/go-helper"

type xCFG struct {
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

// a pdf will have multiple html files
type xRange struct {
	Start int // 1 based, i.e not 0 based -_-
	End   int
}

type xHTMLFile struct {
	LocalPath string
	RemoteURL string
}

type xPdfile struct {
	Serial    int // 1 based
	HTMLFiles []xHTMLFile
	TheRange  xRange
}

func (hf *xHTMLFile) fileBytes() []byte {
	return []byte(hel.FileBytesMust(hf.LocalPath))
}

type xReplace struct {
	Serial int               `json:"serial"`
	Data   map[string]string `json:"data"`
}

func (t *xReplace) FindStr() string {
	var str = ""
	for i := range t.Data {
		str = i
	}
	return str
}

func (t *xReplace) ReplaceStr() string {
	var str = ""
	for i := range t.Data {
		str = t.Data[i]
	}
	return str
}
