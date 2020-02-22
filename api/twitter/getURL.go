package twitter

import (
	"log"
	"net/http"
	"os"
	"publish_it_everywhere/respond"

	"github.com/go-chi/render"

	"github.com/mrjones/oauth"
)

var (
	requestToken *oauth.RequestToken
	consumer     *oauth.Consumer
	url          string
)

func createLoginURL(w http.ResponseWriter, r *http.Request) error {

	consumer = NewTwitterConsumer(os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"))
	var err error
	callbackURL := r.Host + "/api/twitter/callback"
	callbackURL = "http://127.0.0.1:8080/api/twitter/callback" //TODO: Remove this
	requestToken, url, err = consumer.GetRequestTokenAndUrl(callbackURL)
	if err != nil {
		respond.ErrInternalServer(err)
	}
	log.Printf("Request token: %s", requestToken)
	render.Render(w, r, respond.OK(url))
	return nil
}

//NewTwitterConsumer returns the oauth consumer based on the consumer key and consumer secret.
func NewTwitterConsumer(consumerKey, consumerSecret string) *oauth.Consumer {
	return oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)
}
