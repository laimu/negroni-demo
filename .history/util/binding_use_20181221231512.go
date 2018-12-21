package util

import (
	"fmt"
	"net/http"
	"strings"

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

type BookForm struct {
	Name    string
	Price   float64
	Desc	 string
}

type BindHandler struct {
	Cf ContactForm
}

func (h *BindHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestURI := r.RequestURI
	if strings.HasPrefix(requestURI, "/user") {
		if errs := binding.Bind(r, &h.Cf); errs != nil {
			fmt.Println("Data Bind Error!!!")
			http.Error(rw, errs.Error(), http.StatusBadRequest)
			return
		}
	}

	fmt.Println("binding ....")
	next(rw, r)
}
