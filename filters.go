package main

import (
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"strings"
)

var SiteBasedTags = map[string][]string{
	"alorpothe.wordpress.com": {"div#secondary", "div.menu-search", "nav#nav-single",
		"footer#colophon", "div.widget", "div#fb-root", "div#actionbar",
		"form", "h3#reply-title", "div#jp-post-flair", "div#header-img", ".cs-rating .pd-rating",
		"div.wpcnt", "h3#entry-format", ".rating-star-icon", "article.page", "link[rel=dns-prefetch]",
		"#wpcom-block-editor-styles-css",
	},
}
var SpecialReplaceAll = map[string][]string{
	"alorpothe.wordpress.com": {
		"<script src=\"//stats.wp.com/w.js?60\" type=\"text/javascript\" async=\"\" defer=\"\"></script>",
	},
}

func removeTags(htmlBytes []byte) string {

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlBytes))

	if err != nil {
		panic(err)
	}

	tags, ok := SiteBasedTags[DOMAIN]

	if ok {
		for _, tag := range tags {
			doc.Find(tag).Remove()
		}
	}

	htmlStr, err := doc.Html()

	if err != nil {
		panic(err)
	}

	specialReplaces, ok2 := SpecialReplaceAll[DOMAIN]

	if ok2 {
		for _, toBeRemoved := range specialReplaces {
			htmlStr = strings.ReplaceAll(htmlStr, toBeRemoved, "")
		}
	}

	return htmlStr

}
