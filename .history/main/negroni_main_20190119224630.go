package main

import (
	"fmt"
	"net/http"

	util "../util"
	"github.com/jeffbmartinez/delay"
	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	mux.HandleFunc("/book", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello book")
	})

	mux.HandleFunc("/user", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello user")
	})

	n := negroni.Classic() // Includes some default middlewares

	//use gzip
	// n.Use(util.InitGzip())

	bindHandler := new(util.BindHandler)
	n.Use(bindHandler)

	corsHandler := util.InitCors()
	n.Use(corsHandler)

	//add delay, add "X-Add-Delay" to http header
	n.Use(delay.Middleware{})

	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
