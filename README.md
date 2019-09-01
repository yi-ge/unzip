# Unzip

[![Build Status](https://img.shields.io/travis/com/yi-ge/unzip/master.svg)](https://travis-ci.com/yi-ge/unzip)
[![GoDoc](https://godoc.org/github.com/yi-ge/unzip?status.svg)](https://godoc.org/github.com/yi-ge/unzip)

Golang \*.zip decompress.

Fork from [https://github.com/artdarek/go-unzip](https://github.com/artdarek/go-unzip) and remove print, add support for Symlink.

Thank artdarek.

## Usage

```golang
import "github.com/yi-ge/unzip"

u := unzip.New(filePath, outDir)
err := u.Extract()
```

## Notice

由于 Golang1.12 不再支持 Windows XP，因此使用 Golang 自带的 zip 库将在 XP 系统下发生报错，此库对老版本的支持则是自动下载远程`unzip.exe`并自动调用[http://infozip.sourceforge.net/UnZip.html](http://infozip.sourceforge.net/UnZip.html)实现。
