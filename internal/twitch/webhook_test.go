package twitch

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nicklaw5/helix"
)

func TestToBytes(t *testing.T) {
	e := &WebhookEvent{
		Topic:       helix.StreamChangedTopic,
		TopicValues: map[string]string{"user_id": "104137656"},
		Payload:     string([]byte(`{"data":[]}`)),
	}

	b, err := e.ToBytes()
	assert.NoError(t, err)
	assert.Equal(t, []byte(`{"topic":1,"topic_values":{"user_id":"104137656"},"payload":"{\"data\":[]}"}`), b)
}
