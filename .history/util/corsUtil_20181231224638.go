package util

import (
	"github.com/rs/cors"
)

func initCors() cors {
	cors.New(cors.Options{
		AllowedOrigins
	})
}