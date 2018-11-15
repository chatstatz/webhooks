package main

import (
	"sync"
)

// ServiceInterface ...
type ServiceInterface interface {
	Start() error
	Stop()
	PublishMessage([]byte) error
}

// ServiceOptions ...
type ServiceOptions struct {
	Server   ServerInterface
	Producer ProducerInterface
}

// Service ...
type Service struct {
	grWG sync.WaitGroup
	svr  ServerInterface
	pdr  ProducerInterface
}

// newService ...
func newService(options *ServiceOptions) *Service {
	service := &Service{
		svr: options.Server,
		pdr: options.Producer,
	}

	return service
}

// Start starts the service. Start will block.
func (s *Service) Start() error {
	return s.svr.ListenAndServe()
}

// Stop attempts to gracefully stop the service.
func (s *Service) Stop() {
	s.pdr.CloseConn()
}

// PublishMessage ...
func (s *Service) PublishMessage(data []byte) error {
	err := s.pdr.PublishMessage(data)
	if err != nil {
		return err
	}

	return nil
}
