package twitter

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"os"
	"publish_it_everywhere/db"
	"publish_it_everywhere/respond"
	"publish_it_everywhere/schema"
	"publish_it_everywhere/types"
	"publish_it_everywhere/utils"

	"github.com/go-chi/render"

	"github.com/mrjones/oauth"
)

func postStatus(w http.ResponseWriter, r *http.Request) error {
	var postReq types.TwitterPost

	if err := utils.Decode(r, &postReq); err != nil {
		return err
	}
	status := url.PathEscape(postReq.Message)
	if len(status) > 140 {
		respond.ErrBadRequest(errors.New("The number of characters should be less than 140"))
		return nil
	}
	twitterPostURL := "https://api.twitter.com/1.1/statuses/update.json?status=" + status
	consumer := oauth.NewConsumer(
		os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"),
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		},
	)

	var credentials schema.Credentials
	if err := db.FindOne(db.CollectionCredentials, types.JSON{
		"channel_id": postReq.ChannelID,
	}, &credentials); err != nil {
		render.Render(w, r, respond.ErrInternalServer(err))
		return err
	}

	client, err := consumer.MakeHttpClient(&oauth.AccessToken{
		Token:          credentials.Twitter.AccessToken,
		Secret:         credentials.Twitter.AccessSecret,
		AdditionalData: credentials.Twitter.AdditionalData,
	})
	if err != nil {
		render.Render(w, r, respond.ErrInternalServer(err))
		return err
	}
	resp, err := client.Post(twitterPostURL, "application/json", bytes.NewBuffer([]byte{}))
	if err != nil {
		render.Render(w, r, respond.ErrInternalServer(err))
		return err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()
	render.Render(w, r, respond.OK(newStr))
	return nil
}
