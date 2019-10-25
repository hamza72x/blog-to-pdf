package main

import "os"

func bootPaths() {

	buildDir = "./blog-to-pdf-output/" + DOMAIN + "/build"
	urlsTxtPath = "./blog-to-pdf-output/" + DOMAIN + "/urls.txt"
	htmlDir = "./blog-to-pdf-output/" + DOMAIN + "/htmls"
	pdfDir = "./blog-to-pdf-output/" + DOMAIN + "/pdf"

	dirs := []string{buildDir, htmlDir, pdfDir}

	for _, path := range dirs {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err = os.MkdirAll(path, os.ModePerm)
			check(err)
		}
	}
}
