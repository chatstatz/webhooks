package mocks

import (
	m "github.com/stretchr/testify/mock"
)

// Service ...
type Service struct {
	m.Mock
}

// Start ...
func (_m *Service) Start() error {
	args := _m.Called()
	return args.Error(0)
}

// Stop ...
func (_m *Service) Stop() {
	_m.Called()
}

// PublishMessage ...
func (_m *Service) PublishMessage(data []byte) error {
	args := _m.Called(data)
	return args.Error(0)
}
