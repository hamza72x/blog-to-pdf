package main

const UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1"

// go build && ./blog-to-pdf -domain=islamshajid.blogspot.com -force-urls-fetch=true -generate-pdf=false

var siteMapNotApplicables = []string{
	"www.muslimmedia.info",
}

var ArticleWrappers = map[string]string{
	"alorpothe.wordpress.com":   "#content",
	"islamshajid.blogspot.com":  "div.post",
	"amarspondon.wordpress.com": "main#main",
}

var DefaultElementsToBeRemoved = []string{
	"link[rel=dns-prefetch]", "footer", "aside", ".sidebar", ".search", "form",
	".respond", ".wpcnt", ".sharing", ".post-sidebar", ".widget", "#jp-post-flair",
	"#wpcom-block-editor-styles-css",
}
var DefaultStringsToBeReplaced = map[string]string{
	"<script src=\"//stats.wp.com/w.js?60\" type=\"text/javascript\" async=\"\" defer=\"\"></script>": "",
	",v=\"//\"":                                                                                       ",v=\"https://\"",
}

var DivsToBeRemoved = map[string][]string{
	"alorpothe.wordpress.com": {
		"div#secondary", "div.menu-search", "nav#nav-single",
		"footer#colophon", "div#fb-root", "div#actionbar",
		"h3#reply-title", "div#header-img", ".cs-rating .pd-rating",
		"h3#entry-format", ".rating-star-icon", "article.page",
	},
	"islamshajid.blogspot.com": {
		".subscribe-section-container",
		"#PopularPosts1", ".comment-replybox-thread",
	},
	"amarspondon.wordpress.com": {
		".menu-wrapper", ".post-nav-wrapper",
	},
}
var SpecialStringsReplaceAll = map[string][]map[string]string{
	"alorpothe.wordpress.com": {
		{},
	},
}
