package linkedin

import (
	"errors"
	"net/http"
	"publish_it_everywhere/db"
	"publish_it_everywhere/respond"
	"publish_it_everywhere/schema"
	"publish_it_everywhere/types"
	"publish_it_everywhere/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"golang.org/x/oauth2"

	"github.com/go-chi/render"
)

func publisher(w http.ResponseWriter, r *http.Request) error {
	var req types.LinkedinPublishRequest
	if err := utils.Decode(r, &req); err != nil {
		render.Render(w, r, respond.ErrBadRequest(err))
		return err
	}

	credentials := &schema.Credentials{}
	if err := db.FindOne(db.CollectionCredentials, types.JSON{
		"channel_id": req.ChannelID,
	}, credentials); err != nil {
		if err == mongo.ErrNoDocuments {
			render.Render(w, r, respond.ErrBadRequest(errors.New("channel not authenticated")))
		}
		return err
	}

	if time.Now().After(credentials.Linkedin.Expiry) {
		err := errors.New("token expired. please re-authenticate")
		render.Render(w, r, respond.ErrBadRequest(err))
		return err
	}

	token := &oauth2.Token{
		AccessToken: credentials.Linkedin.AccessToken,
		Expiry:      credentials.Linkedin.Expiry,
	}

	client := linkedinConfig.Client(bg, token)
	resp := types.JSON{}
	if err := utils.PostJSON(client, "https://api.linkedin.com/v2/shares", types.JSON{
		"owner": "urn:li:person:" + credentials.Linkedin.UserID,
		"text": types.JSON{
			"text": req.Message,
		},
	}, &resp); err != nil {
		render.Render(w, r, respond.ErrBadRequest(err))
		return err
	}

	render.Render(w, r, respond.OK(types.LinkedinPublishResponse{
		Publisher: credentials.Linkedin.UserName,
		Message:   req.Message,
	}))
	return nil
}
