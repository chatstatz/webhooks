package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nicklaw5/helix"

	"github.com/nicklaw5/chatstatz-webhooks/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/healthz", nil)
	assert.Nil(t, err)

	handler := http.HandlerFunc(healthCheckHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestTwitchWebhookHandler_Get_SubscribeOrUnsubscribe(t *testing.T) {
	challenge := "k63Fl_Fwz3cUk3M1Niicp8zl5pHBCGsJmx7oRPbG"
	subType := "subscribe"
	topic := "https://api.twitch.tv/helix/streams?user_id=104137656"

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/twitch/webhooks?hub.challenge=%s&hub.mode=%s&hub.topic=%s", challenge, subType, topic), nil)
	assert.Nil(t, err)

	handler := http.HandlerFunc(twitchWebhookHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, challenge, rr.Body.String())
}

func TestTwitchWebhookHandler_Get_SubscribeOrUnsubscribe_NoChallenge(t *testing.T) {
	challenge := ""
	subType := "subscribe"
	topic := "https://api.twitch.tv/helix/streams?user_id=104137656"

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/twitch/webhooks?hub.challenge=%s&hub.mode=%s&hub.topic=%s", challenge, subType, topic), nil)
	assert.Nil(t, err)

	handler := http.HandlerFunc(twitchWebhookHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, challenge, rr.Body.String())
}

func TestTwitchWebhookHandler_Post_StreamChanged(t *testing.T) {
	mockBody := []byte(`{"data":[]}`)
	mockLinkHeader := "<https://api.twitch.tv/helix/webhooks/hub>; rel=\"hub\", <https://api.twitch.tv/helix/streams?user_id=104137656>; rel=\"self\""

	mockEvent := &webhookEvent{
		Topic:       helix.StreamChangedTopic,
		TopicValues: map[string]string{"user_id": "104137656"},
		Payload:     string(mockBody),
	}
	mockEventBytes, err := mockEvent.ToBytes()
	assert.Nil(t, err)

	service = &mocks.Service{}
	service.(*mocks.Service).On("PublishMessage", mockEventBytes).Return(nil).Once()

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/twitch/webhooks", bytes.NewBuffer(mockBody))
	assert.Nil(t, err)

	req.Header.Add("Link", mockLinkHeader)

	handler := http.HandlerFunc(twitchWebhookHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.(*mocks.Service).AssertExpectations(t)
}
