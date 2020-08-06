package main

import (
	hel "github.com/thejini3/go-helper"
	"gopkg.in/ini.v1"
)

/*
originalHTMLDir = ./original-html
combinedHTMLDir = ./combined-html
*/
const version = "2.1"

var originalHTMLDir string
var combinedHTMLDir string

var cfgFile *ini.File
var errIni error

var cfgFilePath string
var cfg iniStruct

// var SiteURL string

func main() {
	hel.P("blog-to-pdf, cli version: " + version)
	flags()
	postSuccessBoot()
	build()
	hel.P("blog-to-pdf, cli version: " + version)
}

func postSuccessBoot() {
	loadCfg()
	parseCfgFile()
	modifyPostCfgInit()
	printLoadedCfg()

	bootPaths()
}

func loadCfg() {
	cfgFile, errIni = ini.Load(cfgFilePath)
	hel.PErr("loading ini file", errIni)
}

func parseCfgFile() {
	err := cfgFile.Section("").MapTo(&cfg)
	hel.PErr("mapping ini file, probably bad data!", err)
}

func modifyPostCfgInit() {

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
}

func printLoadedCfg() {
	hel.PS("[CONFIG STARTS]")
	hel.PrettyPrint(&cfg)
	hel.PE("[CONFIG ENDS]")
}
