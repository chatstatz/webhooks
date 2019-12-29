package nats

// IProducerConn ...
type IProducerConn interface {
	Publish(string, []byte) error
	Close()
}

// Producer ...
type Producer struct {
	queue string
	conn  IProducerConn
}

// NewProducer ...
func NewProducer(conn IProducerConn, queue string) *Producer {
	return &Producer{
		conn:  conn,
		queue: queue,
	}
}

// PublishMessage publishes a sigle message to NATS.
func (p *Producer) PublishMessage(data []byte) error {
	err := p.conn.Publish(p.queue, data)
	if err != nil {
		return err
	}

	return nil
}

// CloseConn ...
func (p *Producer) CloseConn() {
	p.conn.Close()
}
