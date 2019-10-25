package main

import (
	"fmt"
)

var protoCol = "https://"

var buildDir string
var pdfDir string
var htmlDir string
var urlsTxtPath string
var SiteURL string
var forceUrlsFetch bool
var SiteMapURL string

func main() {

	bootFlag()
	bootPaths()

	p("IF app stops here, then just run it again!")

	buildAllHTMLS()
}

func check(err error) {
	if err != nil {
		fmt.Println("err: ", err.Error())
		panic(err)
	}
}
