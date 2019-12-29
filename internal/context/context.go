package context

import (
	"github.com/chatstatz/logger"
	"github.com/chatstatz/webhooks/internal"
)

// Context provides commonly-used services.
type Context struct {
	Logger   logger.ILogger
	Producer internal.IProducer
}

// NewContext creates and return a new Context instance.
func NewContext(l logger.ILogger, p internal.IProducer) *Context {
	return &Context{
		Logger:   l,
		Producer: p,
	}
}
