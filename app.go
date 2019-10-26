package main

import (
	"gopkg.in/ini.v1"
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

var cfg *ini.File
var errIni error

var flagIniPath string

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
		// buildAllHTMLS()
		p("IF app stops here, then just 'go' again!")
	}

}

func boot() {

	cfg, errIni = ini.Load(flagIniPath)

	if errIni != nil {
		pp("Error loading ini file: " + errIni.Error())
	}

	bootIni()
	bootPaths()
}
