package util

import "github.com/phyber/negroni-gzip/gzip"

func InitGzip() {
	gzip.Gzip(gzip.DefaultCompression)
}
