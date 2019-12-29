package nats

import (
	"errors"
	"testing"

	"github.com/chatstatz/webhooks/test/mock"
	"github.com/stretchr/testify/assert"
)

func TestNewProducer(t *testing.T) {
	mQueue := "webhooks"
	mConn := &mock.ProducerConn{}
	producer := NewProducer(mConn, mQueue)

	assert.Equal(t, mConn, producer.conn)
	assert.Equal(t, mQueue, producer.queue)
}

func TestProducerPublishMessage(t *testing.T) {
	queueName := "webhooks"
	mockData := []byte("Some test data")

	mProducerConn := &mock.ProducerConn{}
	mProducerConn.On("Publish", queueName, mockData).Return(nil).Once()

	producer := &Producer{
		conn:  mProducerConn,
		queue: queueName,
	}

	producer.PublishMessage(mockData)

	mProducerConn.AssertExpectations(t)
}

func TestProducerPublishMessageError(t *testing.T) {
	queueName := "webhooks"
	mockData := []byte("Some test data")
	errMsg := "Oops something went wrong :("

	mProducerConn := &mock.ProducerConn{}
	mProducerConn.On("Publish", queueName, mockData).Return(errors.New(errMsg)).Once()

	producer := &Producer{
		conn:  mProducerConn,
		queue: queueName,
	}
	err := producer.PublishMessage(mockData)
	assert.EqualError(t, err, errMsg)

	mProducerConn.AssertExpectations(t)
}

func TestProducerCloseConn(t *testing.T) {
	mProducerConn := &mock.ProducerConn{}
	mProducerConn.On("Close").Once()

	producer := &Producer{
		conn: mProducerConn,
	}
	producer.CloseConn()

	mProducerConn.AssertExpectations(t)
}
