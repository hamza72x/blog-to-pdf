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


	urlsTxtFilePath =  "./" + siteDomain + "/" + "urls.txt"
	originalHtmlDir =  "./" + siteDomain + "/original-html"
	combinedHtmlDir =  "./" + siteDomain + "/combined-html"
	pdfDir 			=  "./" + siteDomain + "/pdf"

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
