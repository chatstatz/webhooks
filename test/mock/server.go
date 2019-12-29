package mock

import (
	m "github.com/stretchr/testify/mock"
)

// HTTPServer is a mock server that represents Go's http.Server.
type HTTPServer struct {
	m.Mock
}

// ListenAndServe is the mock equivalent to http.Server.ListenAndServe().
func (_m *HTTPServer) ListenAndServe() error {
	args := _m.Called()
	return args.Error(0)
}
