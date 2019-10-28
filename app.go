package main

import (
	"gopkg.in/ini.v1"
	"strings"
	"encoding/json"
)

/*
flagIniPath 	= ./any_blog.com.ini
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

var flagIniPath string
var cfg IniData
var strReplaces []StringReplace
var SiteURL string

func main() {

	switch bootFlag() {

	case RunModeInit:

		ps("Initialized the ini file!")
		pm("Now run - ")
		pe("$ blog-to-pdf --ini=configs.any_blog.com.ini")

	case RunModeFailed:

		ps("\n+\tWrong instruction given!")
		pe(ConstHelpStr)

	case RunModeGo:
		boot()
		p("IF app stops here, then just run again!")
		buildAllHTMLS()
	}

}

func boot() {

	loadCfg()
	parseCfg()
	changeSomeCfg()
	printLoadedCfg()

	if fileExists(cfg.StringReplacesFile) {

		err := json.Unmarshal(FileDataToByte(cfg.StringReplacesFile), &strReplaces)

		if err != nil {
			ps("Error parsing (" + cfg.StringReplacesFile + ") : " + err.Error())
		}
	}

	bootPaths()
}

func loadCfg() {
	cfgFile, errIni = ini.Load(flagIniPath)

	if errIni != nil {
		pp("Error loading ini file: " + errIni.Error())
	}
}

func parseCfg() {
	err := cfgFile.Section("").MapTo(&cfg)

	if err != nil {
		pp("Error .MapTo(iniData): " + err.Error())
	}
}

func changeSomeCfg() {
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
