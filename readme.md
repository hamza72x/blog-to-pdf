## INSTALL

```
1. Install wkhtmltopdf

- Mac: brew cask install wkhtmltopdf
- Ubuntu: sudo apt install -y wkhtmltopdf
- Windows: Download wkhtmltopdf.exe (https://wkhtmltopdf.org/downloads.html) file & add that file in your System Environment

2. Install Go, cofigure $GOPATH & $GOROOT

3. go get -u github.com/hamza72x/blog-to-pdf
```


## USAGE
```
Usage of blog-to-pdf:

  -c string
    	(required) run the config file, ex: blog-to-pdf -c config.ini
  -d string
    	(required, if -i is passed) initialization directory name, ex: blog-to-pdf -i -d any-blog-name
  -ec
    	print sample config data to console. ex: blog-to-pdf -ec
  -gc
    	create sample config file. ex: blog-to-pdf -gc
  -i	initialize a new directory for new blog, ex: blog-to-pdf -i -d any-blog-name

```

## Default File/Directory structure
```
/<your_project>/config.ini
/<your_project>/custom.css
/<your_project>/string_replaces.json
/<your_project>/pattern_replaces.json
/<your_project>/urls.txt
/<your_project>/original-html/
/<your_project>/combined-html/
/<your_project>/pdf/
```

## Default `config.ini`

To generate pdf after modifying your configs, use `generate_pdf = true`

It's disabled by default, since every blog needs some modification first

```
domain = your_blog.com

# supports multiple sitemap
# comma is important
sitemaps_url = https://your_blog.com/sitemap.xml,

# in case of wp-json type
# comma is important
# sitemaps_url = https://www.muslimmedia.info/wp-json/wp/v2/posts?per_page=50&post_type=post,https://www.muslimmedia.info/wp-json/wp/v2/posts?per_page=50&post_type=page,

get_sitemap_by_wp_json = false

article_per_pdf = 25

url_file = ./urls.txt

# asc or desc, according to sitemap time or date
# only works during url grab

# better try this -
# $ cat urls.txt | sort -u | tee -a sorted.txt

# for blogspot url sorting, use: github.com/hamza72x/sort-blogspot-urls
# urls should start with 'https'
# $ sort-url-by-path-date urls.txt
# Example output format:
# https://x.blogspot.com/2014/11/blog-post.html
# https://x.blogspot.com/2014/11/blog-post_1.html
# https://x.blogspot.com/2014/11/blog-post_2.html

post_order = desc

# Default name: <min_range>-<max_range>_your_blog.com.pdf
# If you set this then: <min_range>-<max_range>_custom.pdf
# pdf_file_name = custom

# use $ instead of ;
browser_user_agent = Mozilla/5.0 (iPhone$ CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1

# This is one of the important value, since we will merge (article_per_pdf) 10 article in a single PDF
# so, which portion of the HTML will be merged in the main Layout? Ex: 'div#content', 'div.post', 'article'
# for "id", use "$" instead of "#"
article_parent_element = .post

# for "id", use "$" instead of "#"
article_title_class = .post h2.entry-title

append_article_url_in_title = true

append_auto_article_number_in_title = true

# for "id", use "$" instead of "#"
elements_to_remove = footer, aside, .respond

# There will be in need of some REPLACES, that's why had to use JSON file,
# Make sure that's valid JSON file

string_replaces_file = string_replaces.json

# Force Re-fetch htmls from server
force_html_fetch = false

# Force Re-fetch urls from sitemap / by wget
force_urls_fetch = true

# -1 => work with all url
limit_urls = -1

# Generate pdf or not, if false then only combined-html files will be created!
generate_pdf = false

# Only generate non generated PDFs
skip_pdf_creation_if_exists_already = false

pdf_output_dir_path = ./pdf

# A0        =>	841 x 1189 mm
# A1        =>	594 x 841 mm
# A2        =>	420 x 594 mm
# A3        =>	297 x 420 mm
# A4        =>	210 x 297 mm, 8.26
# A5        =>	148 x 210 mm
# A6        =>	105 x 148 mm
# A7        =>	74 x 105 mm
# A8        =>	52 x 74 mm
# A9        =>	37 x 52 mm
# B0        =>	1000 x 1414 mm
# B1        =>	707 x 1000 mm
# B10       =>	31 x 44 mm
# B2        =>	500 x 707 mm
# B3        =>	353 x 500 mm
# B4        =>	250 x 353 mm
# B5        =>	176 x 250 mm, 6.93
# B6        =>	125 x 176 mm
# B7        =>	88 x 125 mm
# B8        =>	62 x 88 mm
# B9        =>	33 x 62 mm
# C5E       =>	163 x 229 mm
# Comm10E   =>	105 x 241 mm, U.S. Common 10 Envelope
# Custom    =>	Unknown, or a user defined size.
# DLE       =>	110 x 220 mm
# Executive =>	7.5 x 10 inches, 190.5 x 254 mm
# Folio     =>	210 x 330 mm
# Ledger    =>	431.8 x 279.4 mm
# Legal     =>	8.5 x 14 inches, 215.9 x 355.6 mm
# Letter    =>	8.5 x 11 inches, 215.9 x 279.4 mm
# Tabloid   =>	279.4 x 431.8 mm

pdf_size = A7


# "Landscape" or "Portrait"
pdf_orientation = Portrait


# UI
custom_css_file = custom.css

# Margin / White spaces for pdf (mm)
pdf_margin_top = 3
pdf_margin_left = 3
pdf_margin_right = 3
pdf_margin_bottom = 3
```


## EXAMPLE

```
$ blog-to-pdf init amarspondon
$ cd amarspondon
$ blog-to-pdf config.ini
```

