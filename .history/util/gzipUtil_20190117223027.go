package util

import "github.com/phyber/negroni-gzip/gzip"

func InitGzip() ne {
	
	return gzip.Gzip(gzip.DefaultCompression)
}
