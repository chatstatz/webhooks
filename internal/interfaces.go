package internal

// IServer ...
type IServer interface {
	ListenAndServe() error
}

// IProducer ...
type IProducer interface {
	PublishMessage([]byte) error
	CloseConn()
}

// // IProducerConn ...
// type IProducerConn interface {
// 	Publish(string, []byte) error
// 	Close()
// }

// ServiceInterface ...
type ServiceInterface interface {
	Start() error
	Stop()
	PublishMessage([]byte) error
}
