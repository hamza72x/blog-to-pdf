package main

import (
	hel "github.com/thejini3/go-helper"
	"gopkg.in/ini.v1"
)

/*
originalHTMLDir = ./original-html
combinedHTMLDir = ./combined-html
*/
const version = "2.2"

var originalHTMLDir string
var combinedHTMLDir string

var cfgFile *ini.File
var errIni error

var cfgFilePath string
var cfg xCFG

// var SiteURL string

func main() {
	hel.Pl("blog-to-pdf, cli version: " + version)
	flags()
	loadCfg()
	bootDirPaths()
	build()
	hel.Pl("blog-to-pdf, cli version: " + version)
}

func loadCfg() {

	// load file
	cfgFile, errIni = ini.Load(cfgFilePath)
	hel.PlP("loading ini file", errIni)

	// parse
	err := cfgFile.Section("").MapTo(&cfg)
	hel.PlP("mapping ini file, probably bad data!", err)

	// post changes / fixes
	cfg.ArticleTitleClass = hashifyDollar(cfg.ArticleTitleClass)
	cfg.ArticleParentElement = hashifyDollar(cfg.ArticleParentElement)
	cfg.ElementsToRemove = hashifyDollars(cfg.ElementsToRemove)

	// SiteURL = cfg.Protocol + cfg.Domain

	if len(cfg.BrowserUserAgent) == 0 {
		cfg.BrowserUserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1"
	} else {
		cfg.BrowserUserAgent = collonifyDollar(cfg.BrowserUserAgent)
	}

	if len(cfg.PdfOutputDirPath) == 0 {
		cfg.PdfOutputDirPath = "./pdf"
	}

	if len(cfg.URLFile) == 0 {
		cfg.URLFile = "./urls.txt"
	}

	if len(cfg.PdfFileName) == 0 {
		cfg.PdfFileName = cfg.Domain
	}

	// print loaded cfgs
	hel.Pl("[CONFIG STARTS]")
	hel.PrettyPrint(&cfg)
	hel.Pl("[CONFIG ENDS]")
}
