package context

import (
	"testing"

	"github.com/chatstatz/webhooks/test/mock"

	"github.com/stretchr/testify/assert"
)

func TestNewContext(t *testing.T) {
	mLogger := &mock.Logger{}
	mProducer := &mock.Producer{}

	ctx := NewContext(mLogger, mProducer)

	assert.Equal(t, mLogger, ctx.Logger)
	assert.Equal(t, mProducer, ctx.Producer)

}
