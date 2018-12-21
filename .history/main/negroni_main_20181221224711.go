package main

import (
	"fmt"
	"net/http"

	util "../util"
	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	n := negroni.Classic() // Includes some default middlewares

	bindHandler := new(util.BindHandler)
	n.Use(userHandler)
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
