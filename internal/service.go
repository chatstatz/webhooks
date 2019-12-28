package internal

// Service ...
type Service struct {
	svr IServer
	pdr IProducer
}

// NewService ...
func NewService(server IServer, producer IProducer) *Service {
	service := &Service{
		svr: server,
		pdr: producer,
	}

	return service
}

// Start starts the service. Start will block.
func (s *Service) Start() error {
	return s.svr.ListenAndServe()
}

// Stop attempts to gracefully stop the service. It also
// closes any open connection that remains to the producer.
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
