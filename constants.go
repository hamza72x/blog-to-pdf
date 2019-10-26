package main

const ConstUserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1"


const ConstSampleINI = `protocol = https://

domain = any_blog.com

sitemap_url = any_blog.com/sitemap.xml

article_per_pdf = 10

# This is one of the important value, since we will merge (article_per_pdf) 10 article in a single PDF
# so, which portion of the HTML will be merged in the main Layout? Ex: 'div#content', 'div.post', 'article'
article_parent_element = body

article_title_class = h3.post-title

# for "id", use "$" instead of "#"
elements_to_remove = .comment-form, .post-footer, .post-sidebar, .share-buttons-container, footer, $jp-post-flair, $wpcom-block-editor-styles-css, .wpcnt, .respond, link[rel=dns-prefetch], .back-button-container, .subscribe-section-container, $PopularPosts1, .comment-replybox-thread, .back-button-container, .footer, aside, .sidebar, .search, form,

# There will be in need of some REPLACES, that's why had to use JSON file,
# Make sure that's valid JSON file

string_replaces_file = string_replaces.any_blog.com.json

# Force Re-fetch htmls from server
force_html_fetch = false

# Force Re-fetch urls from sitemap / by wget
force_urls_fetch = false

# if yes, then it will work with only 10 urls
is_test_run = false


# Generate pdf or not, if false then only combined-html files will be created!
generate_pdf = true

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
# font_size = 1.5em
# font_family = "Kohinoor Bangla", "Kalpurush", "Open Sans", serif

# Margin / White spaces for pdf (mm)
pdf_margin_top = 3
pdf_margin_left = 1
pdf_margin_right = 1
pdf_margin_bottom = 3
`

const ConstReplaces = `
[
  {
    "<script src=\"//stats.wp.com/w.js?60\" type=\"text/javascript\" async=\"\" defer=\"\"></script>": ""
  },
  {
    ",v=\"//\"": ",v=\"https://\""
  }
]
`

const ConstHelpStr = `
+
+	# Initialize
+	$ blog-to-pdf init
+
+	# Edit your .ini file according to your need, then -
+	$ blog-to-pdf go --ini=<your_ini_file.ini>
+
`
