


## INSTALL

```
brew cask install wkhtmltopdf
go get gitlab.com/thejini3/blog-to-pdf
cd $GOPATH/src/gitlab.com/thejini3/blog-to-pdf
go install
```
## Check `constants.go` file according to your needs

## USAGE
```
Usage of blog-to-pdf:

  -article-parent-div article
    	Example: article or `div.post`.
    	The parent div of article, specify this, if you want to remove unwanted divs inside the <body> tag! (default "body")

  -article-per-pdf int
    	The number of articles per pdf (default 10)

  -domain string
    	(Required) Domain of the site, Ex: alorpothe.wordpress.com (default "-")

  -force-html-fetch
    	Re-fetch htmls from server if it's not already fetched in local directory

  -force-urls-fetch
    	Re-fetch htmls from server if it's not already fetched in local directory

  -generate-pdf
    	Generate pdf or not, then just html will be created! (default true)
  -https
    	https or not (default true)

  -is-test-run
    	if yes, then it will fetch only 10 url to test!

  -pdf-margin int
    	Margin around the contents of PDF (default 3)

  -pdf-size string
    	The size of output PDF (default "`A7`")

  -sitemap-slug string
    	Sitemap slug, example: sitemap.xml (default "`sitemap.xml`")
```

## EXAMPLE

```
$ blog-to-pdf -domain=alorpothe.wordpress.com -article-per-pdf=7
$ blog-to-pdf -domain=amarspondon.wordpress.com
$ blog-to-pdf -domain=bibijaan.com -sitemap-slug=sitemap-posts.xml -generate-pdf=false -article-parent-div=".inner" -force-html-fetch=true
```

