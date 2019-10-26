#!/usr/bin/env bash
go install
cd storage
blog-to-pdf --ini=configs.islamshajid.blogspot.com.ini
cd ..