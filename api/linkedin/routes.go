package linkedin

import (
	"net/http"
	"publish_it_everywhere/utils"

	"github.com/go-chi/chi"
)

// Init initializes all the search routes
func Init(r chi.Router) {
	r.Method(http.MethodGet, "/authenticate", utils.Handler(authenticator))
	r.Method(http.MethodPost, "/publish", utils.Handler(publisher))
}
