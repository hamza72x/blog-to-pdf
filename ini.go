package main

import (
	"os"
	"strconv"
)

const sampleConfigFileName = "blog_name"

func generateIniFile() string {

	var iniFileName = getConfigFileName(1)

	f, err := os.Create(iniFileName)

	if err != nil {
		pp("Error creating ini file: " + err.Error())
	}

	f2, err2 := os.Create("string_replaces.json")
	f3, _ := os.Create("custom.css")

	if err2 != nil {
		pp("Error creating ini file: " + err2.Error())
	}

	defer f.Close()
	defer f2.Close()
	defer f3.Close()

	f.WriteString(ConstSampleINI)
	f2.WriteString(ConstReplaces)
	f3.WriteString(ConstCusotmCss)

	return iniFileName
}

func getConfigFileName(i int) string {
	if !fileExists(sampleConfigFileName + ".ini") {
		return sampleConfigFileName + ".ini"
	} else if !fileExists(sampleConfigFileName + "_" + strconv.Itoa(i) + ".ini") {
		return sampleConfigFileName + "_" + strconv.Itoa(i) + ".ini"
	}
	return getConfigFileName(i + 1)
}
