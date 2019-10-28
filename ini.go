package main

import (
	"os"
	"strconv"
)

const sampleConfigFileName = "blog_name.ini"

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
	if !fileExists(sampleConfigFileName) {
		return sampleConfigFileName
	} else if !fileExists(sampleConfigFileName + strconv.Itoa(i)) {
		return sampleConfigFileName + strconv.Itoa(i)
	}
	return getConfigFileName(i + 1)
}
