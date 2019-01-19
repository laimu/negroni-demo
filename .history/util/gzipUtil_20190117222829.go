package util

import "github.com/phyber/negroni-gzip/gzip"

func InitGzip() &handler {
	gzip.Gzip(gzip.DefaultCompression)
}
