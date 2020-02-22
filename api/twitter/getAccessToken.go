package twitter

import (
	"fmt"
	"net/http"
	"publish_it_everywhere/respond"

	"github.com/go-chi/render"
)

func getAccessToken(w http.ResponseWriter, r *http.Request) error {
	accessToken, err := consumer.AuthorizeToken(requestToken, r.URL.Query()["oauth_verifier"][0])
	fmt.Println("accessToken:", accessToken)
	if err != nil {
		render.Render(w, r, respond.ErrInternalServer(err))
	}

	render.HTML(w, r, "You have successfully linked your twitter account with publish_it_everywhere")
	return nil
}
