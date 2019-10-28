package main

import (
	"gopkg.in/ini.v1"
	"strings"
	"encoding/json"
)

/*
flagIniPath 	= ./config.ini
urlsTxtFilePath = ./any_blog.com/urls.txt
originalHtmlDir = ./any_blog.com/original-html
combinedHtmlDir = ./any_blog.com/combined-html
pdfDir 			= ./any_blog.com/pdf
 */

var originalHtmlDir string
var combinedHtmlDir string
var pdfDir string

var urlsTxtFilePath string

var cfgFile *ini.File
var errIni error

var iniFilePath string
var cfg IniData
var strReplaces []StringReplace
var SiteURL string

func main() {
	bootFlag()
	successBoot()
	p("IF app stops here, then just run again!")
	buildAllHTMLS()
}

func successBoot() {

	loadCfg()
	parseCfgFile()
	modifyPostCfgInit()
	printLoadedCfg()

	if fileExists(cfg.StringReplacesFile) {
		err := json.Unmarshal(getFileBytes(cfg.StringReplacesFile), &strReplaces)
		pErr("parsing ("+cfg.StringReplacesFile+")", err)
	}

	bootPaths()
}

func loadCfg() {
	cfgFile, errIni = ini.Load(iniFilePath)
	pErr("loading ini file", errIni)
}

func parseCfgFile() {
	err := cfgFile.Section("").MapTo(&cfg)
	pErr("mapping ini file, probably bad data!", err)
}

func modifyPostCfgInit() {

	cfg.ArticleTitleClass = hashifyDollar(cfg.ArticleTitleClass)
	cfg.ArticleParentElement = hashifyDollar(cfg.ArticleParentElement)
	cfg.ElementsToRemove = hashifyDollars(cfg.ElementsToRemove)

	SiteURL = cfg.Protocol + cfg.Domain

	if !strings.Contains(cfg.SiteMapURL, "https://") || !strings.Contains(cfg.SiteMapURL, "http://") {
		cfg.SiteMapURL = cfg.Protocol + cfg.SiteMapURL
	}
}
func printLoadedCfg() {
	ps("[CONFIG STARTS]")
	PrettyPrint(&cfg)
	pe("[CONFIG ENDS]")
}
