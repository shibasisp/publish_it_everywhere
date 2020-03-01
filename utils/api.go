package utils

import (
	"net/http"

	"github.com/gorilla/context"
	log "github.com/sirupsen/logrus"
)

// Handler custom api handler help us to handle all the errors in one place
type Handler func(w http.ResponseWriter, r *http.Request) error

func (f Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := f(w, r)
	defer context.Clear(r)
	if err != nil {
		log.Errorf("Error: %s\n", err.Error())
	}
}
