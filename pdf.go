package main

import (
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"bytes"
	"io/ioutil"
	"log"
)

func pdfFromLocalFile(htmlFile HtmlFile) {

	pdfg := wkhtmltopdf.NewPDFPreparer()
	htmlfile, err := ioutil.ReadFile(htmlFile.LocalPath)
	if err != nil {
		log.Fatal(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(bytes.NewReader(htmlfile)))
	pdfg.Dpi.Set(600)

	// The contents of htmlsimple.html are saved as base64 string in the JSON file
	jb, err := pdfg.ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	// Server code
	pdfgFromJSON, err := wkhtmltopdf.NewPDFGeneratorFromJSON(bytes.NewReader(jb))
	if err != nil {
		log.Fatal(err)
	}

	err = pdfgFromJSON.Create()
	if err != nil {
		log.Fatal(err)
	}


	fmt.Printf("PDF size %vkB\n", len(pdfg.Bytes())/1024)
}

func pdfFromURL(htmlFile HtmlFile) {

	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		panic(err)
	}

	pdfg.Dpi.Set(600)
	pdfg.NoCollate.Set(false)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginBottom.Set(40)
	pdfg.MarginLeft.Set(0)

	page1 := wkhtmltopdf.NewPage(htmlFile.URL)

	page1.DisableSmartShrinking.Set(true)
	page1.HeaderSpacing.Set(10.01)
	page1.Allow.Set("/usr/local/html")
	page1.Allow.Set("/usr/local/images")
	page1.CustomHeader.Set("User-Agent", UserAgent)
	page1.ViewportSize.Set("1024x1024")

	pdfg.AddPage(page1)

	// pdfg.Cover.Input = "https://wkhtmltopdf.org/index.html"
	pdfg.Cover.Zoom.Set(0.75)

	pdfg.TOC.Include = true
	pdfg.TOC.DisableDottedLines.Set(true)

	err = pdfg.Create()

	if err != nil {
		panic(err)
	}

	err = pdfg.WriteFile(htmlFile.LocalPath)

	if err != nil {
		panic(err)
	}

	fmt.Printf("PDF size %vkB\n", len(pdfg.Bytes())/1024)
}
