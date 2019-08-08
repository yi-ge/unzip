# Unzip

Golang \*.zip decompress.

[![Build Status](https://img.shields.io/travis/com/yi-ge/unzip/master.svg)](https://travis-ci.com/yi-ge/unzip)
[![GoDoc](https://godoc.org/github.com/yi-ge/unzip?status.svg)](https://godoc.org/github.com/yi-ge/unzip)

Golang \*.tar.xz decompress.

Fork from [https://github.com/artdarek/go-unzip](https://github.com/artdarek/go-unzip) and remove print, add support for Symlink.

Thank artdarek.

## Usage

```golang
import "github.com/yi-ge/unzip"

u := unzip.New(filePath, outDir)
err := u.Extract()
```
