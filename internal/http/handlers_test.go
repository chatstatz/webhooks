package http

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chatstatz/webhooks/internal/context"
	"github.com/chatstatz/webhooks/internal/twitch"
	"github.com/chatstatz/webhooks/test/mock"

	"github.com/nicklaw5/helix"
	"github.com/stretchr/testify/assert"
	m "github.com/stretchr/testify/mock"
)

func TestHealthCheckHandler(t *testing.T) {
	server := &Server{}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/healthz", nil)
	assert.Nil(t, err)

	handler := http.HandlerFunc(server.healthCheckHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestTwitchWebhookHandlerGetSubscribeOrUnsubscribe(t *testing.T) {
	mLogger := &mock.Logger{}
	mLogger.On("Debugf", m.Anything).Times(2)
	mProducer := &mock.Producer{}

	server := &Server{
		ctx: context.NewContext(mLogger, mProducer),
	}

	challenge := "k63Fl_Fwz3cUk3M1Niicp8zl5pHBCGsJmx7oRPbG"
	subType := "subscribe"
	topic := "https://api.twitch.tv/helix/streams?user_id=104137656"

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/twitch/webhooks?hub.challenge=%s&hub.mode=%s&hub.topic=%s", challenge, subType, topic), nil)
	assert.Nil(t, err)

	handler := http.HandlerFunc(server.twitchWebhookHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, challenge, rr.Body.String())

	mLogger.AssertExpectations(t)
}

func TestTwitchWebhookHandlerGetSubscribeOrUnsubscribe_NoChallenge(t *testing.T) {
	mLogger := &mock.Logger{}
	mLogger.On("Debugf", m.Anything).Times(1)
	mProducer := &mock.Producer{}

	server := &Server{
		ctx: context.NewContext(mLogger, mProducer),
	}

	challenge := ""
	subType := "subscribe"
	topic := "https://api.twitch.tv/helix/streams?user_id=104137656"

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/twitch/webhooks?hub.challenge=%s&hub.mode=%s&hub.topic=%s", challenge, subType, topic), nil)
	assert.Nil(t, err)

	handler := http.HandlerFunc(server.twitchWebhookHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, challenge, rr.Body.String())

	mLogger.AssertExpectations(t)
}

func TestTwitchWebhookHandler_Post_StreamChanged(t *testing.T) {
	mockBody := []byte(`{"data":[]}`)
	mockLinkHeader := "<https://api.twitch.tv/helix/webhooks/hub>; rel=\"hub\", <https://api.twitch.tv/helix/streams?user_id=104137656>; rel=\"self\""

	mockEvent := &twitch.WebhookEvent{
		Topic:       helix.StreamChangedTopic,
		TopicValues: map[string]string{"user_id": "104137656"},
		Payload:     string(mockBody),
	}
	mockEventBytes, err := mockEvent.ToBytes()
	assert.Nil(t, err)

	mLogger := &mock.Logger{}
	mLogger.On("Debugf", m.Anything).Times(5)
	mProducer := &mock.Producer{}
	mProducer.On("PublishMessage", mockEventBytes).Return(nil).Once()

	server := &Server{
		ctx: context.NewContext(mLogger, mProducer),
	}

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/twitch/webhooks", bytes.NewBuffer(mockBody))
	assert.Nil(t, err)

	req.Header.Add("Link", mockLinkHeader)

	handler := http.HandlerFunc(server.twitchWebhookHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	mLogger.AssertExpectations(t)
}
