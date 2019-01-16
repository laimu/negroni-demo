package util

import (
	"github.com/rs/cors"
)

func initCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"http://foo.com"},
	})
}
