package main

import "os"

func bootPaths() {

	buildDir = "./storage/" + DOMAIN + "/build"
	urlsTxtPath = "./storage/" + DOMAIN + "/urls.txt"
	htmlDir = "./storage/" + DOMAIN + "/htmls"
	pdfDir = "./storage/" + DOMAIN + "/pdf"

	dirs := []string{buildDir, htmlDir, pdfDir}

	for _, path := range dirs {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err = os.MkdirAll(path, os.ModePerm)
			check(err)
		}
	}
}
