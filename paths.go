package main

import "os"

const OutputDir = "./blog-to-pdf-output"

func bootPaths() {

	buildDir = OutputDir + "/" + DOMAIN + "/build"
	urlsTxtPath = OutputDir + "/" + DOMAIN + "/urls.txt"
	htmlDir = OutputDir + "/" + DOMAIN + "/htmls"
	pdfDir = OutputDir + "/" + DOMAIN + "/pdf"

	dirs := []string{buildDir, htmlDir, pdfDir}

	for _, path := range dirs {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err = os.MkdirAll(path, os.ModePerm)
			check(err)
		}
	}
}
