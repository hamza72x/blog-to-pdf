package main

import "os"

/*
flagIniPath 	= ./blog_name.ini
urlsTxtFilePath = ./urls.txt
originalHtmlDir = ./original-html
combinedHtmlDir = ./combined-html
pdfDir 			= ./pdf
 */

func bootPaths() {

	urlsTxtFilePath = "./urls.txt"
	originalHtmlDir = "./original-html"
	combinedHtmlDir = "./combined-html"
	pdfDir = "./pdf"

	createDirsIfNotExists(
		[]string{originalHtmlDir, combinedHtmlDir, pdfDir},
	)
}

func createDirsIfNotExists(dirs []string) {
	for _, path := range dirs {
		if !pathExists(path) {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				p("Error creating directory: " + err.Error())
			}
		}
	}
}
