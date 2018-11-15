package main

// ProducerInterface ...
type ProducerInterface interface {
	PublishMessage([]byte) error
	CloseConn()
}

// ProducerConnInterface ...
type ProducerConnInterface interface {
	Publish(string, []byte) error
	Close()
}

// Producer ...
type Producer struct {
	host  string
	queue string
	conn  ProducerConnInterface
}

type newProducerCallback func(host, queue string) (*Producer, error)

func newProducer(host, queue string, cb newProducerCallback) (*Producer, error) {
	return cb(host, queue)
}

// PublishMessage ...
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
