package schema

// TwitterOAuth stores the mapping between channelID and RequestToken
type TwitterOAuth struct {
	ChannelID    string       `json:"channel_id,omitempty" bson:"channel_id,omitempty"`
	RequestToken RequestToken `json:"request_token,omitempty" bson:"request_token,omitempty"`
}

// RequestToken stores the OAuth token and OAuth verifier.
type RequestToken struct {
	OAuthToken    string `json:"oauth_token" bson:"oauth_token"`
	OAuthVerifier string `json:"oauth_verifier" bson:"oauth_verifier"`
}
