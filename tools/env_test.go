package tools

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	tests := []struct {
		key   string
		value string
	}{
		{"FOO", "bar"},
	}

	for _, test := range tests {
		err := os.Setenv(test.key, test.value)
		assert.NoError(t, err)

		val := GetEnv(test.key, "none")
		assert.Equal(t, test.value, val)

		err = os.Unsetenv(test.key)
		assert.NoError(t, err)
	}
}

func TestGetEnvFallback(t *testing.T) {
	tests := []struct {
		key      string
		fallback string
	}{
		{"FOO", "bar"},
	}

	for _, test := range tests {
		val := GetEnv(test.key, test.fallback)
		assert.Equal(t, test.fallback, val)
	}
}
