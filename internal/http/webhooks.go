package http

import (
	"encoding/json"

	"github.com/nicklaw5/helix"
)

// WebhookEvent ...
type WebhookEvent struct {
	Topic       helix.WebhookTopic `json:"topic"`
	TopicValues map[string]string  `json:"topic_values"`
	Payload     string             `json:"payload"`
}

// ToBytes ...
func (we *WebhookEvent) ToBytes() ([]byte, error) {
	return json.Marshal(we)
}
