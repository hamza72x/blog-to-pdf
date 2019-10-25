package main

import (
	"fmt"
	"unicode/utf8"
	"io/ioutil"
	"strings"
	"os"
	"net/http"
	"time"
)

func getURLContent(urlStr string) []byte {

	// fmt.Printf("HTML code of %s ...\n", urlStr)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create and modify HTTP request before sending
	request, err := http.NewRequest("GET", urlStr, nil)

	if err != nil {
		panic(err)
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1")

	// Make request
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	htmlBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return htmlBytes
}

func getFileContents(filePath string) []byte {

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error reading file: " + filePath)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("Error ioutil.ReadAll: " + filePath)
	}

	return b
}
func fileExists(filename string) bool {

	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func removeSpecialChars(str string) string {
	str = strings.ReplaceAll(str, "/", "-")
	str = strings.ReplaceAll(str, ":", "-")
	return str
}

func checkDomain(name string) error {

	switch {
	case len(name) == 0:
		return nil // an empty domain name will result in a cookie without a domain restriction
	case len(name) > 255:
		return fmt.Errorf("cookie domain: name length is %d, can't exceed 255", len(name))
	}
	var l int
	for i := 0; i < len(name); i++ {
		b := name[i]
		if b == '.' {
			// check domain labels validity
			switch {
			case i == l:
				return fmt.Errorf("cookie domain: invalid character '%c' at offset %d: label can't begin with a period", b, i)
			case i-l > 63:
				return fmt.Errorf("cookie domain: byte length of label '%s' is %d, can't exceed 63", name[l:i], i-l)
			case name[l] == '-':
				return fmt.Errorf("cookie domain: label '%s' at offset %d begins with a hyphen", name[l:i], l)
			case name[i-1] == '-':
				return fmt.Errorf("cookie domain: label '%s' at offset %d ends with a hyphen", name[l:i], l)
			}
			l = i + 1
			continue
		}
		// test label character validity, note: tests are ordered by decreasing validity frequency
		if !(b >= 'a' && b <= 'z' || b >= '0' && b <= '9' || b == '-' || b >= 'A' && b <= 'Z') {
			// show the printable unicode character starting at byte offset i
			c, _ := utf8.DecodeRuneInString(name[i:])
			if c == utf8.RuneError {
				return fmt.Errorf("cookie domain: invalid rune at offset %d", i)
			}
			return fmt.Errorf("cookie domain: invalid character '%c' at offset %d", c, i)
		}
	}
	// check top level domain validity
	switch {
	case l == len(name):
		return fmt.Errorf("cookie domain: missing top level domain, domain can't end with a period")
	case len(name)-l > 63:
		return fmt.Errorf("cookie domain: byte length of top level domain '%s' is %d, can't exceed 63", name[l:], len(name)-l)
	case name[l] == '-':
		return fmt.Errorf("cookie domain: top level domain '%s' at offset %d begins with a hyphen", name[l:], l)
	case name[len(name)-1] == '-':
		return fmt.Errorf("cookie domain: top level domain '%s' at offset %d ends with a hyphen", name[l:], l)
	case name[l] >= '0' && name[l] <= '9':
		return fmt.Errorf("cookie domain: top level domain '%s' at offset %d begins with a digit", name[l:], l)
	}
	return nil
}

func p(str string) {
	fmt.Println("=======================================================")
	fmt.Println(str)
	fmt.Println("=======================================================")
}

func ContainsStr(array []string, value string) bool {
	for _, a := range array {
		if a == value {
			return true
		}
	}
	return false
}

func tempWrite(path string, str string) {
	f, _ := os.Create(path)
	f.WriteString(str)
	f.Close()
}
