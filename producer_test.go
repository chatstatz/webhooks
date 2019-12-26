package main

import (
	"errors"
	"testing"

	"github.com/chatstatz/webhooks/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewProducer(t *testing.T) {
	mockHost := "http://blah.blah"
	mockQueue := "webhooks"
	mockConn := &mocks.ProducerConn{}

	mockProducer := &Producer{
		conn:  mockConn,
		host:  mockHost,
		queue: mockQueue,
	}

	cb := func(host, queue string) (*Producer, error) {
		return &Producer{
			conn:  &mocks.ProducerConn{},
			host:  host,
			queue: queue,
		}, nil
	}

	producer, err := newProducer(mockHost, mockQueue, cb)

	assert.NoError(t, err)
	assert.Equal(t, mockProducer, producer)
}

func TestNewProducerError(t *testing.T) {
	errMsg := "Faield to connect to nats producer"

	cb := func(host, queue string) (*Producer, error) {
		return nil, errors.New(errMsg)
	}

	producer, err := newProducer("", "", cb)

	assert.Nil(t, producer)
	assert.EqualError(t, err, errMsg)
}

func TestProducerPublishMessage(t *testing.T) {
	queueName := "webhooks"
	mockData := []byte("Some test data")

	mockProducerConn := new(mocks.ProducerConn)
	mockProducerConn.On("Publish", queueName, mockData).Return(nil).Once()

	producer := &Producer{
		conn:  mockProducerConn,
		queue: queueName,
	}
	producer.PublishMessage(mockData)

	mockProducerConn.AssertExpectations(t)
}

func TestProducerPublishMessageError(t *testing.T) {
	queueName := "webhooks"
	mockData := []byte("Some test data")
	errMsg := "Oops something went wrong :("

	mockProducerConn := new(mocks.ProducerConn)
	mockProducerConn.On("Publish", queueName, mockData).Return(errors.New(errMsg)).Once()

	producer := &Producer{
		conn:  mockProducerConn,
		queue: queueName,
	}
	err := producer.PublishMessage(mockData)
	assert.EqualError(t, err, errMsg)

	mockProducerConn.AssertExpectations(t)
}

func TestProducerCloseConn(t *testing.T) {
	mockProducerConn := new(mocks.ProducerConn)
	mockProducerConn.On("Close").Once()

	producer := &Producer{
		conn: mockProducerConn,
	}
	producer.CloseConn()

	mockProducerConn.AssertExpectations(t)
}
