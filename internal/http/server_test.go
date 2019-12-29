package http

import (
	"fmt"
	"net"
	"testing"

	"github.com/chatstatz/webhooks/internal/context"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	mockHost := "127.0.0.0"
	mockPort := "9999"

	server := NewServer(&context.Context{}, mockHost, mockPort)

	fmt.Println(server.Addr)

	assert.Equal(t, server.Addr, net.JoinHostPort(mockHost, mockPort))
}
