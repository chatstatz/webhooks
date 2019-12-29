package mock

import (
	m "github.com/stretchr/testify/mock"
)

// Producer is a mock NATS producer.
type Producer struct {
	m.Mock
}

// PublishMessage is the mock equivalent to Producer.PublishMessage().
func (_m *Producer) PublishMessage(data []byte) error {
	args := _m.Called(data)
	return args.Error(0)
}

// CloseConn is the mock equivalent to Producer.CloseConn().
func (_m *Producer) CloseConn() {
	_m.Called()
}

// ProducerConn is a mock NATS TCP connection.
type ProducerConn struct {
	m.Mock
}

// Publish is the mock equivalent to ProducerConn.Publish().
func (_m *ProducerConn) Publish(queue string, data []byte) error {
	args := _m.Called(queue, data)
	return args.Error(0)
}

// Close is the mock equivalent to ProducerConn.Close().
func (_m *ProducerConn) Close() {
	_m.Called()
}
