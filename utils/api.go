package utils

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/gorilla/context"
	"google.golang.org/appengine/log"
)

// Handler custom api handler help us to handle all the errors in one place
type Handler func(w http.ResponseWriter, r *http.Request) error

func (f Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := f(w, r)
	defer context.Clear(r)
	if err != nil {
		log.Errorf(r.Context(), "Error: %s\n", err.Error())
		// respond.Fail(w, err)
		render.JSON(w, r, err)
	}
}
