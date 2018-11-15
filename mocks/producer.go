package mocks

import (
	m "github.com/stretchr/testify/mock"
)

// Producer ...
type Producer struct {
	m.Mock
}

// PublishMessage ...
func (_m *Producer) PublishMessage(data []byte) error {
	args := _m.Called(data)
	return args.Error(0)
}

// CloseConn ...
func (_m *Producer) CloseConn() {
	_m.Called()
}

// ProducerConn ...
type ProducerConn struct {
	m.Mock
}

// Publish ...
func (_m *ProducerConn) Publish(queue string, data []byte) error {
	args := _m.Called(queue, data)
	return args.Error(0)
}

// Close ...
func (_m *ProducerConn) Close() {
	_m.Called()
}
