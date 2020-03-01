package types

// LinkedinPublishRequest ...
type LinkedinPublishRequest struct {
	ChannelID string `json:"channel_id" bson:"channel_id"`
	Message   string `json:"message" bson:"message"`
}
