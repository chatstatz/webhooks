package nats

import (
	"time"

	"github.com/nats-io/nats.go"
)

// // IProducer ...
// type IProducer interface {
// 	PublishMessage([]byte) error
// 	CloseConn()
// }

// IProducerConn ...
type IProducerConn interface {
	Publish(string, []byte) error
	Close()
}

// Producer ...
type Producer struct {
	host  string
	queue string
	conn  IProducerConn
}

// NewProducer ...
func NewProducer(host, queue string) (*Producer, error) {
	conn, err := nats.Connect(host, nats.Timeout(3*time.Second))
	if err != nil {
		return nil, err
	}

	return &Producer{
		conn:  conn,
		host:  host,
		queue: queue,
	}, nil
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
