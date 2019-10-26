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

	if flag.Arg(0) == "init" {

		generateIniFile()

		return RunModeInit

	} else {

		if len(flagIniPath) <= 1 {
			pp("Please specify --ini=<your_ini_file.ini>")
		}

		if !fileExists(flagIniPath) {
			ps(flagIniPath + " doesn't exist!")
			pm("To create ini file, run  -")
			pe("$ blog-to-pdf init")
			os.Exit(0)
		}

		return RunModeGo
	}

	return RunModeFailed

}
