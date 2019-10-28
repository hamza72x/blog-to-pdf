package main

import (
	"flag"
	"os"
)

func bootFlag() RunMode {

	flag.StringVar(&flagIniPath, "ini", "-",
		"(REQUIRED) Pass -ini to execute that config!",
	)

	flag.Parse()

	if len(flag.Args()) == 0 {
		return RunModeFailed
	}

	if flag.Arg(0) == "init" {

		generateIniFile()

		return RunModeInit

	}

	if !fileExists(flagIniPath) {
		ps(flagIniPath + " doesn't exist!")
		pm("To create ini file, run  -")
		pe("$ blog-to-pdf init")
		os.Exit(0)
	}

	return RunModeGo
}
