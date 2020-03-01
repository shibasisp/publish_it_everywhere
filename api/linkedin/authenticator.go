package linkedin

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"publish_it_everywhere/db"
	"publish_it_everywhere/respond"
	"publish_it_everywhere/schema"
	"publish_it_everywhere/types"
	"strings"

	"golang.org/x/oauth2"

	"github.com/go-chi/render"
)

func authenticator(w http.ResponseWriter, r *http.Request) error {
	code, channelID, err := getCodeAndChannelID(r)
	if err != nil {
		render.Render(w, r, respond.ErrBadRequest(err))
		return err
	}

	tok, err := linkedinConfig.Exchange(bg, code)
	if err != nil {
		render.Render(w, r, respond.ErrBadRequest(err))
		return err
	}

	userID, userName, err := getUserDetails(tok)
	if err != nil {
		render.Render(w, r, respond.ErrInternalServer(err))
		return err
	}

	credentials := &schema.Credentials{}
	if err := db.Update(db.CollectionCredentials, types.JSON{
		"channel_id": channelID,
	}, types.JSON{
		"$set": types.JSON{
			"linkedin": schema.LinkedinCredentials{
				AccessToken: tok.AccessToken,
				Expiry:      tok.Expiry,
				UserID:      userID,
				UserName:    userName,
			},
		},
	}, credentials, true); err != nil {
		render.Render(w, r, respond.ErrInternalServer(err))
		return err
	}

	render.HTML(w, r, "You have successfully linked your linkedin account with publish_it_everywhere. Close this tab.")
	return nil
}

func getCodeAndChannelID(r *http.Request) (string, string, error) {
	keys, ok := r.URL.Query()["code"]
	if !ok || len(keys) < 1 {
		return "", "", errors.New("code not passed in query string")
	}
	code := keys[0]

	keys, ok = r.URL.Query()["state"]
	if !ok || len(keys) < 1 {
		return code, "", errors.New("state not passed in query string")
	}
	channelID := keys[0]

	return code, channelID, nil
}

func getUserDetails(token *oauth2.Token) (string, string, error) {
	client := linkedinConfig.Client(bg, token)
	resp, err := client.Get("https://api.linkedin.com/v2/me")
	if err != nil {
		return "", "", err
	}

	bf := new(bytes.Buffer)
	bf.ReadFrom(resp.Body)

	jsonData := types.JSON{}
	if err := json.Unmarshal(bf.Bytes(), &jsonData); err != nil {
		return "", "", err
	}

	id, ok := jsonData["id"].(string)
	if !ok {
		return "", "", errors.New("response was expected to have field: `id`")
	}

	firstName, ok := jsonData["localizedFirstName"].(string)
	if !ok {
		return "", "", errors.New("response was expected to have field: `localizedFirstName`")
	}

	lastName, ok := jsonData["localizedLastName"].(string)
	if !ok {
		return "", "", errors.New("response was expected to have field: `localizedLastName`")
	}

	return id, strings.Join([]string{firstName, lastName}, " "), nil
}
