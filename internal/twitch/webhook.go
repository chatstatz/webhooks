package twitch

import (
	"encoding/json"

	"github.com/nicklaw5/helix"
)

// WebhookEvent represents a Twitch webhook event payload.
// See https://dev.twitch.tv/docs/api/webhooks-reference.
type WebhookEvent struct {
	Topic       helix.WebhookTopic `json:"topic"`
	TopicValues map[string]string  `json:"topic_values"`
	Payload     string             `json:"payload"`
}

// ToBytes converts the json representation of WebhookEvent
// to bytes.
func (we *WebhookEvent) ToBytes() ([]byte, error) {
	return json.Marshal(we)
}
