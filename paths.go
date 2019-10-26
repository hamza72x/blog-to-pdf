package main

import "os"

/*
flagIniPath 	= ./any_blog.com.ini
urlsTxtFilePath = ./any_blog.com/urls.txt
originalHtmlDir = ./any_blog.com/original-html
combinedHtmlDir = ./any_blog.com/combined-html
pdfDir 			= ./any_blog.com/pdf
 */

func bootPaths() {


	urlsTxtFilePath =  "./" + cfg.Domain + "/" + "urls.txt"
	originalHtmlDir =  "./" + cfg.Domain + "/original-html"
	combinedHtmlDir =  "./" + cfg.Domain + "/combined-html"
	pdfDir 			=  "./" + cfg.Domain + "/pdf"

	createDirs(
		[]string{originalHtmlDir, combinedHtmlDir, pdfDir},
	)
}

func createDirs(dirs []string) {
	for _, path := range dirs {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				p("Error creating directory: " + err.Error())
			}
		}
	}
}
