// Package logger provides logging utilities
package logger

// Logger defines the logger contract
type Logger interface {
	// Info logs a message at Info levevl.
	// Args must be provided as key-value pairs
	Info(msg string, args ...any)

	// Debug logs a message at Debug levevl.
	// Args must be provided as key-value pairs
	Debug(msg string, args ...any)

	// Warn logs a message at Warn levevl.
	// Args must be provided as key-value pairs
	Warn(msg string, args ...any)

	// Error logs a message at Error levevl.
	// Args must be provided as key-value pairs
	Error(msg string, args ...any)

	// Return new logger with context(static key-value pairs)
	// Args must be a provided as key-value pairs
	With(msg string, args ...any) Logger
}
