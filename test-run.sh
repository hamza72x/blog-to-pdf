#!/usr/bin/env bash
go install
cd storage/amarspondhon
# blog-to-pdf init amarspondhon
blog-to-pdf config.ini
cd ..
cd ..