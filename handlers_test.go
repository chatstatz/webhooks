package main

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nicklaw5/chatstatz-webhooks/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/health-check", nil)
	assert.Nil(t, err)

	handler := http.HandlerFunc(healthCheckHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, `{"success":true}`, rr.Body.String())
}

func TestHealthCheckHandlerInvalidMethod(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/health-check", nil)
	assert.Nil(t, err)

	handler := http.HandlerFunc(healthCheckHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
	assert.Equal(t, `{"success":false,"message":"method not allowed"}`, rr.Body.String())
}

func TestTwitchWebhookHandler(t *testing.T) {
	mockBody := []byte(`{"test":"data"}`)

	service = &mocks.Service{}
	service.(*mocks.Service).On("PublishMessage", mockBody).Return(nil).Once()

	bufBody := bytes.NewBuffer(mockBody)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/twitch/webhooks", bufBody)
	assert.Nil(t, err)

	handler := http.HandlerFunc(twitchWebhookHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	service.(*mocks.Service).AssertExpectations(t)
}

func TestTwitchWebhookHandlerBadMehtod(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/twitch/webhooks", nil)
	assert.Nil(t, err)

	handler := http.HandlerFunc(twitchWebhookHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
}

func TestTwitchWebhookHandlerNilBody(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/twitch/webhooks", nil)
	assert.Nil(t, err)

	handler := http.HandlerFunc(twitchWebhookHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestTwitchWebhookHandlerFailedToPublishMessage(t *testing.T) {
	mockBody := []byte(`{"test":"data"}`)
	mockError := errors.New("bad things happened")

	service = &mocks.Service{}
	service.(*mocks.Service).On("PublishMessage", mockBody).Return(mockError).Once()

	bufBody := bytes.NewBuffer(mockBody)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/twitch/webhooks", bufBody)
	assert.Nil(t, err)

	handler := http.HandlerFunc(twitchWebhookHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	service.(*mocks.Service).AssertExpectations(t)
}
