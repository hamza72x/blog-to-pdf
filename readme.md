## INSTALL

```
brew cask install wkhtmltopdf
go get gitlab.com/thejini3/blog-to-pdf
cd $GOPATH/src/gitlab.com/thejini3/blog-to-pdf
go install
```


## USAGE
```
Usage of blog-to-pdf:

  $ blog-to-pdf init <dir_name>
  
  # Edit auto-generated .ini file, then -

  $ blog-to-pdf <config_file.ini>

```

## EXAMPLE

```
$ blog-to-pdf init amarspondon
$ cd amarspondon
$ blog-to-pdf amarspondon.ini
```

