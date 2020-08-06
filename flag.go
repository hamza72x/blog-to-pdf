package main

import (
	"flag"
	"fmt"
	"os"

	hel "github.com/thejini3/go-helper"
)

func flags() {

	init := flag.Bool("i", false, "initialize a new directory for new blog, ex: blog-to-pdf -i -d any-blog-name")
	dir := flag.String("d", "", "(required, if -i is passed) initialization directory name, ex: blog-to-pdf -i -d any-blog-name")
	echoConfig := flag.Bool("ec", false, "print sample config data to console. ex: blog-to-pdf -ec")
	generateIni := flag.Bool("gc", false, "create sample config file. ex: blog-to-pdf -gc")

	flag.StringVar(&cfgFilePath, "c", "", "(required) run the config file, ex: blog-to-pdf -c config.ini")

	flag.Parse()

	if *init {

		if *dir == "" {
			hel.Pl("Err: -d is required during -i")
			flag.PrintDefaults()
			os.Exit(1)
		}

		if hel.PathExists(*dir) {
			panic("Dir `" + *dir + "` already exists, use different name!")
		}

		handleFlagInit(*dir)
		os.Exit(0)
	}

	if *echoConfig {
		fmt.Println(constSampleINI)
		os.Exit(0)
	}

	if *generateIni {
		fname := hel.GetNonCreatedFileName("config", ".ini", 1)
		if err := hel.StrToFile(fname, constSampleINI); err == nil {
			hel.Pl("Generated: " + fname)
		}
		os.Exit(0)
	}

	if cfgFilePath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if !hel.FileExists(cfgFilePath) {
		hel.PS("The ini file `" + cfgFilePath + "` doesn't exist!")
		hel.PM("To auto-generate ini file, run  -")
		hel.PE("$ blog-to-pdf -i -d any-blog-name")
		os.Exit(1)
	}

}

func handleFlagInit(dir string) {

	dir = hel.AZ_AND_NUMBER_ONLY(dir)

	createDirsIfNotExists([]string{dir})

	hel.P("Created directory: " + dir)

	var filesAndData = map[string]string{
		dir + "/" + "config.ini":            constSampleINI,
		dir + "/" + "string_replaces.json":  constReplacesJSONStr,
		dir + "/" + "pattern_replaces.json": constReplacesJSONPatternStr,
		dir + "/" + "custom.css":            constCusotmCSS,
	}

	for filename, fileData := range filesAndData {
		hel.StrToFile(filename, fileData)
	}

	var instruction = `
+		Now -
+		$ cd ` + dir + `
+		Edit config.ini according to your needs, then -
+		$ blog-to-pdf -c config.ini
`
	hel.P(instruction)
}
