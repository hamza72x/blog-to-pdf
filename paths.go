package main

import (
	"os"

	hel "github.com/hamza02x/go-helper"
)

/*
flagIniPath 	= ./blog_name.ini
originalHTMLDir = ./original-html
combinedHTMLDir = ./combined-html
*/

func bootDirPaths() {

	originalHTMLDir = cfgDir + "/original-html"
	combinedHTMLDir = cfgDir + "/combined-html"

	createDirsIfNotExists(
		[]string{originalHTMLDir, combinedHTMLDir, cfg.PdfOutputDirPath},
	)
}

func createDirsIfNotExists(dirs []string) {
	for _, path := range dirs {
		if !hel.PathExists(path) {
			err := os.MkdirAll(path, os.ModePerm)
			hel.PlP("creating directory: ", err)
		}
	}
}
