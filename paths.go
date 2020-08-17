package main

import (
	"os"

	hel "github.com/thejini3/go-helper"
)

/*
flagIniPath 	= ./blog_name.ini
originalHTMLDir = ./original-html
combinedHTMLDir = ./combined-html
*/

func bootDirPaths() {

	originalHTMLDir = "./original-html"
	combinedHTMLDir = "./combined-html"

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
