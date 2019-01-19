package util

import (
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"
)

func InitGzip() negroni.Handler {

	return gzip.Gzip(gzip.DefaultCompression)
}
