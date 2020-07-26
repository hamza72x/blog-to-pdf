package main

import (
	"flag"
	"fmt"
	"os"

	hel "github.com/thejini3/go-helper"
)

func handleBootFlags() {

	flag.Parse()

	if len(flag.Args()) == 0 {
		runFailed()
	}

	/// - init
	if flag.Arg(0) == "init" {

		dir := flag.Arg(1)

		if len(dir) == 0 {
			runFailed()
		} else if hel.PathExists(dir) {
			hel.OSExit("Dir `" + dir + "` already exists, use different name!")
		}

		handleFlagInit(dir)
		os.Exit(0)
	} else if flag.Arg(0) == "echo-config.ini" {

		fmt.Println(constSampleINI)

		os.Exit(0)
	}

	/// - generate-ini
	if flag.Arg(0) == "generate-ini" {
		createFile(hel.GetNonCreatedFileName("config", ".ini", 1), constSampleINI)
		os.Exit(0)
	}

	iniFilePath = flag.Arg(0)

	if !hel.FileExists(iniFilePath) {
		hel.PS("The ini file `" + iniFilePath + "` doesn't exist!")
		hel.PM("To auto-generate ini file, run  -")
		hel.PE("$ blog-to-pdf init")
		os.Exit(0)
	}

}

func runFailed() {
	hel.PS("\n+\tWrong instruction given!")
	hel.PE(constHelpStr)
	hel.OSExit("")
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
		createFile(filename, fileData)
	}

	var instruction = `
+		Now -
+		$ cd ` + dir + `
+		Edit config.ini according to your needs, then -
+		$ blog-to-pdf config.ini
`
	hel.P(instruction)
}

func createFile(filename string, fileData string) {

	file, err := os.Create(filename)

	if err != nil {
		hel.OSExit("Error creating `" + filename + "` file: " + err.Error())
	}

	if err != nil {
		hel.OSExit("Error creating `" + filename + "` file: " + err.Error())
	}

	hel.P("Created string replace file: " + filename)

	file.WriteString(fileData)

	file.Close()
}
