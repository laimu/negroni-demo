package util

import (
	"github.com/rs/cors"
)

func InitCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"http://foo.com"},
	})
}
