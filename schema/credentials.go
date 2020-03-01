package schema

import "time"

//Credentials stores all the credentials along with channelID
type Credentials struct {
	ChannelID string              `json:"channel_id,omitempty" bson:"channel_id,omitempty"`
	Twitter   TwitterCredentials  `json:"twitter,omitempty" bson:"twitter,omitempty"`
	Linkedin  LinkedinCredentials `json:"linkedin,omitempty" bson:"linkedin,omitempty"`
}

// TwitterCredentials stores the twitter credentials
type TwitterCredentials struct {
	AccessToken    string            `json:"access_token,omitempty" bson:"access_token,omitempty"`
	AccessSecret   string            `json:"access_secret,omitempty" bson:"access_secret,omitempty"`
	AdditionalData map[string]string `json:"additional_data,omitempty" bson:"additional_data,omitempty"`
	RequestToken   RequestToken      `json:"request_token,omitempty" bson:"request_token,omitempty"`
}

// LinkedinCredentials ...
type LinkedinCredentials struct {
	AccessToken string    `json:"access_token,omitempty" bson:"access_token,omitempty"`
	Expiry      time.Time `json:"expiry,omitempty" bson:"expiry,omitempty"`
	UserName    string    `json:"user_name,omitempty" bson:"user_name,omitempty"`
	UserID      string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
}
