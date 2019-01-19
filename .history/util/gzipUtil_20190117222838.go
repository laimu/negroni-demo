package util

import "github.com/phyber/negroni-gzip/gzip"

func InitGzip() handler {
	return &gzip.Gzip(gzip.DefaultCompression)
}
