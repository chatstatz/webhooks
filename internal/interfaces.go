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

// ServiceInterface ...
type ServiceInterface interface {
	Start() error
	Stop()
	PublishMessage([]byte) error
}
