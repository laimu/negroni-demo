package main

import (
	"fmt"
	"net/http"

	"github.com/mholt/binding"
	"github.com/urfave/negroni"
)

type UserHandler struct {
	Cf ContactForm
}

func (h *UserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if errs := binding.Bind(r, &h.Cf); errs != nil {
		fmt.Println("Data Bind Error!!!")
		rw.Write([]byte("Data Bind Error!!!"))
		return
	}
	next(rw, r)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	n := negroni.Classic() // Includes some default middlewares

	userHandler := new(UserHandler)
	n.UseHandler(mux)
	n.With(userHandler)

	http.ListenAndServe(":3000", n)
}
