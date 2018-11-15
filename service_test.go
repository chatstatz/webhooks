package main

import (
	"errors"
	"net"
	"strconv"
	"testing"

	"github.com/nicklaw5/chatstatz-webhooks/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	mockHost := "127.0.0.0"
	mockPort := 9999

	server := newServer(mockHost, mockPort)

	assert.Equal(t, server.Addr, net.JoinHostPort(mockHost, strconv.Itoa(mockPort)))
}

func TestServiceStart(t *testing.T) {
	mockHTTPServer := new(mocks.HTTPServer)
	mockHTTPServer.On("ListenAndServe").Return(nil).Once()

	service := newService(&ServiceOptions{
		Server: mockHTTPServer,
	})
	service.Start()

	mockHTTPServer.AssertExpectations(t)
}

func TestServiceStop(t *testing.T) {
	mockProducer := new(mocks.Producer)
	mockProducer.On("CloseConn").Once()

	service := newService(&ServiceOptions{
		Producer: mockProducer,
	})
	service.Stop()

	mockProducer.AssertExpectations(t)
}

func TestServicePublishMessage(t *testing.T) {
	mockData := []byte("Some test data")

	mockProducer := new(mocks.Producer)
	mockProducer.On("PublishMessage", mockData).Return(nil).Once()

	service := newService(&ServiceOptions{
		Producer: mockProducer,
	})
	service.PublishMessage(mockData)

	mockProducer.AssertExpectations(t)
}

func TestServicePublishMessageError(t *testing.T) {
	errMsg := "Oops something went wrong :("
	mockData := []byte("Some test data")

	mockProducer := new(mocks.Producer)
	mockProducer.On("PublishMessage", mockData).Return(errors.New(errMsg)).Once()

	service := newService(&ServiceOptions{
		Producer: mockProducer,
	})
	err := service.PublishMessage(mockData)

	assert.EqualError(t, err, errMsg)
	mockProducer.AssertExpectations(t)
}
