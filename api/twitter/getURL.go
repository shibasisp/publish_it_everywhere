package twitter

import (
	"log"
	"net/http"
	"os"
	"publish_it_everywhere/db"
	"publish_it_everywhere/respond"
	"publish_it_everywhere/schema"

	"github.com/go-chi/render"

	"github.com/mrjones/oauth"
)

var (
	consumer *oauth.Consumer
)

func createLoginURL(w http.ResponseWriter, r *http.Request) error {

	consumer = NewTwitterConsumer(os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"))
	callbackURL := r.Host + "/api/twitter/callback"
	callbackURL = "http://127.0.0.1:8080/api/twitter/callback" //TODO: Remove this
	rqstToken, url, err := consumer.GetRequestTokenAndUrl(callbackURL)
	if err != nil {
		respond.ErrInternalServer(err)
		return nil
	}
	twitterOAuth := schema.TwitterOAuth{
		ChannelID: r.URL.Query()["channel_id"][0],
		RequestToken: schema.RequestToken{
			OAuthToken:    rqstToken.Token,
			OAuthVerifier: rqstToken.Secret,
		},
	}
	if err := db.Insert(db.CollectionTwitterOAuth, twitterOAuth); err != nil {
		log.Printf("err in insertion %v", err)
		respond.ErrInternalServer(err)
		return nil
	}

	log.Printf("Request token: %s", rqstToken)
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
