package main

import "os"

func bootPaths() {

	buildDir = "./blog-to-pdf/" + DOMAIN + "/build"
	urlsTxtPath = "./blog-to-pdf/" + DOMAIN + "/urls.txt"
	htmlDir = "./blog-to-pdf/" + DOMAIN + "/htmls"
	pdfDir = "./blog-to-pdf/" + DOMAIN + "/pdf"

	dirs := []string{buildDir, htmlDir, pdfDir}

	for _, path := range dirs {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err = os.MkdirAll(path, os.ModePerm)
			check(err)
		}
	}
}
