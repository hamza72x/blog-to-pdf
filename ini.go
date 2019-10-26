package main

import (
	"os"
)

func generateIniFile() {

	f, err := os.Create("configs.any_blog.com.ini")

	if err != nil {
		pp("Error creating ini file: " + err.Error())
	}
	defer f.Close()

	f2, err2 := os.Create("string_replaces.any_blog.com.json")

	if err2 != nil {
		pp("Error creating ini file: " + err2.Error())
	}

	defer f2.Close()

	f2.WriteString(ConstReplaces)
	f.WriteString(ConstSampleINI)
}
