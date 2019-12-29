package internal

// WebhooksService ...
type WebhooksService struct {
	svr IServer
	pdr IProducer
}

// NewWebhooksService ...
func NewWebhooksService(server IServer, producer IProducer) *WebhooksService {
	return &WebhooksService{
		svr: server,
		pdr: producer,
	}
}

// Start starts the service. Start will block.
func (s *WebhooksService) Start() error {
	return s.svr.ListenAndServe()
}

// Stop attempts to gracefully stop the service. It also
// closes any open connection that remains to the producer.
func (s *WebhooksService) Stop() {
	s.pdr.CloseConn()
}

// PublishMessage ...
func (s *WebhooksService) PublishMessage(data []byte) error {
	err := s.pdr.PublishMessage(data)
	if err != nil {
		return err
	}

	return nil
}
