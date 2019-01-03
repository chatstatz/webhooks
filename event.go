package main

import (
	"encoding/json"

	"github.com/nicklaw5/helix"
)

type webhookEvent struct {
	Topic       helix.WebhookTopic `json:"topic"`
	TopicValues map[string]string  `json:"topic_values"`
	Payload     string             `json:"payload"`
}

func (we *webhookEvent) ToBytes() ([]byte, error) {
	return json.Marshal(we)
}
