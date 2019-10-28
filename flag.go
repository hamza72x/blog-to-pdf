package main

import (
	"flag"
	"os"
)

func bootFlag() (RunMode, string) {

	flag.Parse()

	if len(flag.Args()) == 0 {
		return RunModeFailed, ""
	}

	if flag.Arg(0) == "init" {
		return RunModeInit, generateIniFile()
	}

	flagIniPath = flag.Arg(0)

	if !fileExists(flagIniPath) {
		ps("The ini file `" + flagIniPath + "` doesn't exist!")
		pm("To auto-generate ini file, run  -")
		pe("$ blog-to-pdf init")
		os.Exit(0)
	}

	return RunModeGo, ""
}
