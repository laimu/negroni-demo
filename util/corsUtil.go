package util

import (
	"fmt"

	"github.com/rs/cors"
)

func InitCors() *cors.Cors {
	fmt.Println("cors..")
	return cors.New(cors.Options{
		AllowedOrigins: []string{"http://foo.com"},
	})
}
