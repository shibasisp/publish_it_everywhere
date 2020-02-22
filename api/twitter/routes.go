package twitter

import (
	"net/http"
	"publish_it_everywhere/utils"

	"github.com/go-chi/chi"
)

// Init initializes all the search routes
func Init(r chi.Router) {
	r.Method(http.MethodGet, "/authurl", utils.Handler(createLoginURL))
	r.Method(http.MethodGet, "/callback", utils.Handler(getAccessToken))
}
