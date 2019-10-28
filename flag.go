package main

import (
	"flag"
	"os"
)

func bootFlag() {

	flag.Parse()

	if len(flag.Args()) == 0 {
		runFailed()
	}

	if flag.Arg(0) == "init" {

		dir := flag.Arg(1)

		if len(dir) == 0 {
			runFailed()
		} else if pathExists(dir) {
			pp("Dir `" + dir + "` already exists, use different name!")
		}

		handleFlagInit(dir)
		os.Exit(0)
	}

	iniFilePath = flag.Arg(0)

	if !fileExists(iniFilePath) {
		ps("The ini file `" + iniFilePath + "` doesn't exist!")
		pm("To auto-generate ini file, run  -")
		pe("$ blog-to-pdf init")
		os.Exit(0)
	}

}

func runFailed() {
	ps("\n+\tWrong instruction given!")
	pe(ConstHelpStr)
	pp("")
}
