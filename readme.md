## INSTALL

```
go get gitlab.com/thejini3/blog-to-pdf
cd $GOPATH/src/gitlab.com/thejini3/blog-to-pdf
go install
```

## USAGE
```
Usage of blog-to-pdf:
  -article-per-pdf int
    	The number of articles per pdf (default 10)

  -domain string
    	(Required) Domain of the site, Ex: alorpothe.wordpress.com

  -force-html-fetch
    	Re-fetch htmls from server if it's not already fetched in local directory

  -force-sitemap-fetch
    	Re-fetch htmls from server if it's not already fetched in local directory

  -https
    	https or not (default true)
```

## EXAMPLE
```
blog-to-pdf -domain=alorpothe.wordpress.com -article-per-pdf=7
```