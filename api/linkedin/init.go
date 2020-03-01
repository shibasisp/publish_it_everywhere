package linkedin

import (
	"context"
	"net/url"

	"publish_it_everywhere/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/linkedin"
)

var bg context.Context
var linkedinConfig *oauth2.Config

// Initialize ...
func Initialize() {
	bg = context.Background()

	u, _ := url.Parse(config.SelfURL)
	u.Path = "/api/linkedin/authenticate"

	linkedinConfig = &oauth2.Config{
		ClientID:     config.LinkedinClientID,
		ClientSecret: config.LinkedinClientSecret,
		RedirectURL:  u.String(),
		Endpoint:     linkedin.Endpoint,
	}
}
