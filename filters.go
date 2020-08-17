package main

import (
	"encoding/json"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
	hel "github.com/thejini3/go-helper"
)

func additionalFilter(str string) string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(str))

	if err != nil {
		panic(err)
	}

	for _, tag := range cfg.ElementsToRemove {
		doc.Find(tag).Remove()
	}

	htmlStr, err := doc.Html()

	if err != nil {
		panic(err)
	}

	htmlStr = strings.ReplaceAll(htmlStr, lineBreakHTML3, "")
	htmlStr = strings.ReplaceAll(htmlStr, lineBreakHTML4, "")

	if patternReplaces := getFileJSONIfExists(cfg.PatternReplacesFile); patternReplaces != nil {
		for _, replace := range patternReplaces {
			rgx := regexp.MustCompile(replace.FindStr())
			rgx.ReplaceAllStringFunc(htmlStr, func(str string) string {
				// hel.Pl("[Pattern]", k, "Replacing", str, "with", `$1`)
				htmlStr = strings.ReplaceAll(htmlStr, rgx.ReplaceAllString(str, `$1`), replace.ReplaceStr())
				return ""
			})
		}
	}

	if strReplaces := getFileJSONIfExists(cfg.StringReplacesFile); strReplaces != nil {
		for _, replace := range strReplaces {
			// hel.Pl(replace.FindStr(), "with", replace.ReplaceStr(), strings.Contains(htmlStr, replace.FindStr()))
			htmlStr = strings.ReplaceAll(htmlStr, replace.FindStr(), replace.ReplaceStr())
		}
	}

	return htmlStr

}

const lineBreakHTML4 = `
<br/>
<br/>
<br/>
<br/>
`
const lineBreakHTML3 = `
<br/>
<br/>
<br/>
`

func getFileJSONIfExists(filename string) []xReplace {
	var v []xReplace

	if hel.FileExists(filename) {
		err := json.Unmarshal(hel.FileBytesMust(filename), &v)
		if err != nil {
			panic("parsing (" + filename + ") - " + err.Error())
		}
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i].Serial < v[j].Serial
	})
	return v
}
