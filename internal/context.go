package internal

import (
	"github.com/chatstatz/logger"
)

// Context ...
type Context struct {
	Logger   logger.ILogger
	Producer IProducer
}
