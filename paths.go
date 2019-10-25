package main

import "os"

func bootPaths() {

	buildDir = "./" + DOMAIN + "/build"
	siteMapFilePath = "./" + DOMAIN + "/sitemap.txt"
	htmlDir = "./" + DOMAIN + "/htmls"

	dirs := []string{buildDir, htmlDir}

	for _, path := range dirs {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err = os.MkdirAll(path, os.ModePerm)
			check(err)
		}
	}
}
