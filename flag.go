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

		dir := flag.Arg(1)

		if len(dir) == 0 {
			return RunModeFailed
		} else if pathExists(dir) {
			pp("Dir `" + dir + "` already exists, use different name!")
		}

		handleFlagInit(dir)
		os.Exit(0)
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
