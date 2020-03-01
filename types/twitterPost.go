package types

//TwitterPost stores the request
type TwitterPost struct {
	Message   string `json:"message"`
	ChannelID string `json:"channel_id"`
}
