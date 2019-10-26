package main

type IniData struct {
	Protocol             string   `ini:"protocol"`
	Domain               string   `ini:"domain"`
	SiteMapURL           string   `ini:"sitemap_url"`
	ArticlePerPDF        int      `ini:"article_per_pdf"`
	ArticleParentElement string   `ini:"article_parent_element"`
	ArticleTitleClass    string   `ini:"article_title_class"`
	ElementsToRemove     []string `ini:"elements_to_remove"`
	StringReplacesFile   string   `ini:"string_replaces_file"`
	ForceFetchHtml       bool     `ini:"force_html_fetch"`
	ForceUrlsFetch       bool     `ini:"force_urls_fetch"`
	IsTestRun            bool     `ini:"is_test_run"`
	GeneratePDF          bool     `ini:"generate_pdf"`
	PdfPageSize          string   `ini:"pdf_size"`
	PdfOrientation       string   `ini:"pdf_orientation"`
	FontSize             string   `ini:"font_size"`
	FontFamily           []string `ini:"font_family"`
	PdfMarginTop         int      `ini:"pdf_margin_top"`
	PdfMarginLeft        int      `ini:"pdf_margin_left"`
	PdfMarginRight       int      `ini:"pdf_margin_right"`
	PdfMarginBottom      int      `ini:"pdf_margin_bottom"`
}

type StringReplace struct {
	ReplaceKey string
	ReplaceVal string
}
