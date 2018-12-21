package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
	"util"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	n := negroni.Classic() // Includes some default middlewares

	userHandler := new(util.UserHandler)
	n.UseHandler(mux)
	n.With(userHandler)

	http.ListenAndServe(":3000", n)
}
