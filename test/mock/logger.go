package mock

import (
	m "github.com/stretchr/testify/mock"
)

// See https://github.com/chatstatz/logger

// Logger is a mock server.
type Logger struct {
	m.Mock
}

// Debug is the mock equivalent to logger.Debug().
func (_m *Logger) Debug(v ...interface{}) {
	_m.Called()
}

// Debugf is the mock equivalent to logger.Debugf().
func (_m *Logger) Debugf(f string, v ...interface{}) {
	_m.Called()
}

// Info is the mock equivalent to logger.Info().
func (_m *Logger) Info(v ...interface{}) {
	_m.Called()
}

// Infof is the mock equivalent to logger.Infof().
func (_m *Logger) Infof(f string, v ...interface{}) {
	_m.Called()
}

// Warn is the mock equivalent to logger.Warn().
func (_m *Logger) Warn(v ...interface{}) {
	_m.Called()
}

// Warnf is the mock equivalent to logger.Warnf().
func (_m *Logger) Warnf(f string, v ...interface{}) {
	_m.Called()
}

// Error is the mock equivalent to logger.Error().
func (_m *Logger) Error(v ...interface{}) {
	_m.Called()
}

// Errorf is the mock equivalent to logger.Errorf().
func (_m *Logger) Errorf(f string, v ...interface{}) {
	_m.Called()
}

// Fatal is the mock equivalent to logger.Fatal().
func (_m *Logger) Fatal(v ...interface{}) {
	_m.Called()
}

// Fatalf is the mock equivalent to logger.Fatalf().
func (_m *Logger) Fatalf(f string, v ...interface{}) {
	_m.Called()
}
