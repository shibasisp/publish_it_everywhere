package twitter

import (
	"log"
	"net/http"
	"publish_it_everywhere/db"
	"publish_it_everywhere/respond"
	"publish_it_everywhere/schema"
	"publish_it_everywhere/types"

	"github.com/go-chi/render"
	"github.com/mrjones/oauth"
)

func getAccessToken(w http.ResponseWriter, r *http.Request) error {
	oauthToken := r.URL.Query()["oauth_token"][0]
	var twitterOAuth schema.TwitterOAuth
	if err := db.FindOne(db.CollectionTwitterOAuth, types.JSON{
		"request_token.oauth_token": oauthToken,
	}, &twitterOAuth); err != nil {
		respond.ErrInternalServer(err)
		return nil
	}
	oAuthRequestToken := oauth.RequestToken{
		Token:  twitterOAuth.RequestToken.OAuthToken,
		Secret: twitterOAuth.RequestToken.OAuthVerifier,
	}
	accessToken, err := consumer.AuthorizeToken(&oAuthRequestToken, r.URL.Query()["oauth_verifier"][0])
	if err != nil {
		log.Println(err)
		render.Render(w, r, respond.ErrInternalServer(err))
		return nil
	}
	credentials := schema.Credentials{
		ChannelID: twitterOAuth.ChannelID,
		Twitter: schema.TwitterCredentials{
			AccessToken:    accessToken.Token,
			AccessSecret:   accessToken.Secret,
			AdditionalData: accessToken.AdditionalData,
			RequestToken:   twitterOAuth.RequestToken,
		},
	}
	if err = db.Update(db.CollectionCredentials, types.JSON{
		"channel_id": credentials.ChannelID,
	}, types.JSON{
		"$set": types.JSON{
			"twitter": credentials.Twitter,
		},
	}, &credentials); err != nil {
		log.Println(err)
		render.Render(w, r, respond.ErrInternalServer(err))
		return nil
	}

	render.HTML(w, r, "You have successfully linked your twitter account with publish_it_everywhere")
	return nil
}
