package util

import "github.com/phyber/negroni-gzip/gzip"

func InitGzip() gzip.handler {
	
	return gzip.Gzip(gzip.DefaultCompression)
}
