package internal

import (
	"errors"
	"net"
	"testing"

	"github.com/chatstatz/webhooks/internal/http"

	"github.com/chatstatz/webhooks/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	mockHost := "127.0.0.0"
	mockPort := "9999"

	server := http.New(mockHost, mockPort)

	assert.Equal(t, server.Addr, net.JoinHostPort(mockHost, mockPort))
}

func TestServiceStart(t *testing.T) {
	mockHTTPServer := new(mocks.HTTPServer)
	mockHTTPServer.On("ListenAndServe").Return(nil).Once()

	service := NewService(mockHTTPServer, nil)
	service.Start()

	mockHTTPServer.AssertExpectations(t)
}

func TestServiceStop(t *testing.T) {
	mockProducer := new(mocks.Producer)
	mockProducer.On("CloseConn").Once()

	service := NewService(nil, mockProducer)
	service.Stop()

	mockProducer.AssertExpectations(t)
}

func TestServicePublishMessage(t *testing.T) {
	mockData := []byte("Some test data")

	mockProducer := new(mocks.Producer)
	mockProducer.On("PublishMessage", mockData).Return(nil).Once()

	service := NewService(nil, mockProducer)
	service.PublishMessage(mockData)

	mockProducer.AssertExpectations(t)
}

func TestServicePublishMessageError(t *testing.T) {
	errMsg := "Oops something went wrong :("
	mockData := []byte("Some test data")

	mockProducer := new(mocks.Producer)
	mockProducer.On("PublishMessage", mockData).Return(errors.New(errMsg)).Once()

	service := NewService(nil, mockProducer)
	err := service.PublishMessage(mockData)

	assert.EqualError(t, err, errMsg)
	mockProducer.AssertExpectations(t)
}
