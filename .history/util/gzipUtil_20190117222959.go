package util

import "github.com/phyber/negroni-gzip/gzip"

func InitGzip() gzip.handler {
	gzip.
	return gzip.Gzip(gzip.DefaultCompression)
}
