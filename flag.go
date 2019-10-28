package main

import (
	"flag"
	"os"
)

func bootFlag() RunMode {

	flag.Parse()

	if len(flag.Args()) == 0 {
		return RunModeFailed
	}

	if flag.Arg(0) == "init" {

		generateIniFile()

		return RunModeInit

	}

	flagIniPath = flag.Arg(0)

	if !fileExists(flagIniPath) {
		ps("The ini file `" + flagIniPath + "` doesn't exist!")
		pm("To auto-generate ini file, run  -")
		pe("$ blog-to-pdf init")
		os.Exit(0)
	}

	return RunModeGo
}
