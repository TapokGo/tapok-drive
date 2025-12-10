// Package logger provides utilities for logging info
package logger

type slog struct{}

// NewSlog creates slog logger
func NewSlog(logPath string) (*slog, error) {
	return &slog{}, nil
}

// Error logs a message at Error level
func (s *slog) Error(msg string, args ...any) {}

// Info logs a message at Info level
func (s *slog) Info(msg string, args ...any) {
	print(msg)
}

// Debug logs a message at Debug level
func (s *slog) Debug(msg string, args ...any) {}

// Warn logs a message at Warn level
func (s *slog) Warn(msg string, args ...any) {}

// With returns new logger with extra key-value pair
func (s *slog) With(rgs ...any) Logger {
	return &slog{}
}
