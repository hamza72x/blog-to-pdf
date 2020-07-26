#!/usr/bin/env bash
go install
#cd storage/muslimskeptic
cd examples/iDream4Life
# blog-to-pdf init amarspondhon
blog-to-pdf config.docx.ini
cd ..
cd ..