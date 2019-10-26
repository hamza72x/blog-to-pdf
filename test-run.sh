#!/usr/bin/env bash
go install
cd storage
blog-to-pdf --ini=dev.ini
cd ..