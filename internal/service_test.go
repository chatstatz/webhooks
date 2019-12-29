package internal

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chatstatz/webhooks/test/mock"
)

func TestNewService(t *testing.T) {
	mServer := &mock.HTTPServer{}
	mProducer := &mock.Producer{}

	s := NewWebhooksService(mServer, mProducer)

	assert.Equal(t, mServer, s.svr)
	assert.Equal(t, mProducer, s.pdr)
}

func TestServiceStart(t *testing.T) {
	mHTTPServer := &mock.HTTPServer{}
	mHTTPServer.On("ListenAndServe").Return(nil).Once()

	s := NewWebhooksService(mHTTPServer, nil)
	s.Start()

	mHTTPServer.AssertExpectations(t)
}

func TestServiceStop(t *testing.T) {
	mProducer := &mock.Producer{}
	mProducer.On("CloseConn").Once()

	s := NewWebhooksService(nil, mProducer)
	s.Stop()

	mProducer.AssertExpectations(t)
}

func TestServicePublishMessage(t *testing.T) {
	mockData := []byte("Some test data")

	mProducer := &mock.Producer{}
	mProducer.On("PublishMessage", mockData).Return(nil).Once()

	service := NewWebhooksService(nil, mProducer)
	service.PublishMessage(mockData)

	mProducer.AssertExpectations(t)
}

func TestServicePublishMessageError(t *testing.T) {
	errMsg := "Oops something went wrong :("
	mockData := []byte("Some test data")

	mProducer := &mock.Producer{}
	mProducer.On("PublishMessage", mockData).Return(errors.New(errMsg)).Once()

	service := NewWebhooksService(nil, mProducer)
	err := service.PublishMessage(mockData)

	assert.EqualError(t, err, errMsg)
	mProducer.AssertExpectations(t)
}
