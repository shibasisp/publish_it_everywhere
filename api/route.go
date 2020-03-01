package api

import (
	"publish_it_everywhere/api/linkedin"
	"publish_it_everywhere/api/twitter"

	"github.com/go-chi/chi"
)

// Init initializes all the search routes
func Init(r chi.Router) {
	r.Route("/twitter", twitter.Init)
	r.Route("/linkedin", linkedin.Init)
}
