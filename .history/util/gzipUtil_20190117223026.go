package util

import "github.com/phyber/negroni-gzip/gzip"

func InitGzip()  {
	
	return gzip.Gzip(gzip.DefaultCompression)
}
