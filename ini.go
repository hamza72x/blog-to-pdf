package main

import (
	"os"
)

func generateIniFile() {

	f, err := os.Create("configs.any_blog.com.ini")

	if err != nil {
		pp("Error creating ini file: " + err.Error())
	}

	f2, err2 := os.Create("string_replaces.any_blog.com.json")
	f3, _ := os.Create("custom.css")

	if err2 != nil {
		pp("Error creating ini file: " + err2.Error())
	}

	defer f.Close()
	defer f2.Close()
	defer f3.Close()

	f2.WriteString(ConstReplaces)
	f.WriteString(ConstSampleINI)
}
