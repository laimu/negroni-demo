package util

import (
	"fmt"
	"net/http"

	"github.com/mholt/binding"
)

type ContactForm struct {
	User struct {
		ID int
	}
	Email   string
	Message string
}

// Then provide a field mapping (pointer receiver is vital)
func (cf *ContactForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&cf.User.ID: "user_id",
		&cf.Email:   "email",
		&cf.Message: binding.Field{
			Form:     "message",
			Required: true,
		},
	}
}

// Now your handlers can stay clean and simple
func handler(resp http.ResponseWriter, req *http.Request) {
	contactForm := new(ContactForm)
	if errs := binding.Bind(req, contactForm); errs != nil {
		http.Error(resp, errs.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(resp, "From:    %d\n", contactForm.User.ID)
	fmt.Fprintf(resp, "Message: %s\n", contactForm.Message)
}

type UserHandler struct {
	Cf ContactForm
}

func (h *UserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println(r.RequestURI)
	if errs := binding.Bind(r, &h.Cf); errs != nil {
		fmt.Println("Data Bind Error!!!")
		http.Error(rw, errs.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("binding ....")
	next(rw, r)
}
