package main

const UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1"

// go build && ./blog-to-pdf -domain=islamshajid.blogspot.com -force-urls-fetch=true -generate-pdf=false

var siteMapSites = []string{
	"alorpothe.wordpress.com",
	"islamshajid.blogspot.com",
	"amarspondon.wordpress.com",
}

var ArticleWrappers = map[string]string{
	"alorpothe.wordpress.com":   "#content",
	"islamshajid.blogspot.com":  "div.post",
	"amarspondon.wordpress.com": "main#main",
}

var DivsToBeRemoved = map[string][]string{
	"alorpothe.wordpress.com": {"div#secondary", "div.menu-search", "nav#nav-single",
		"footer#colophon", "div.widget", "div#fb-root", "div#actionbar",
		"form", "h3#reply-title", "div#jp-post-flair", "div#header-img", ".cs-rating .pd-rating",
		"div.wpcnt", "h3#entry-format", ".rating-star-icon", "article.page", "link[rel=dns-prefetch]",
		"#wpcom-block-editor-styles-css",
	},
	"islamshajid.blogspot.com": {
		"aside", ".post-sidebar", ".subscribe-section-container", ".search",
		"#PopularPosts1", "footer", ".comment-replybox-thread", ".sharing",
	},
	"amarspondon.wordpress.com": {
		".menu-wrapper", ".respond", "aside", "footer", ".wpcnt", ".post-nav-wrapper", "#jp-post-flair",
	},
}
var SpecialStringsReplaceAll = map[string][]map[string]string{
	"alorpothe.wordpress.com": {
		{"<script src=\"//stats.wp.com/w.js?60\" type=\"text/javascript\" async=\"\" defer=\"\"></script>": ""},
		{",v=\"//\"": ",v=\"https://\""},
	},
}
