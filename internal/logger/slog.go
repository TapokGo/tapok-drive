// Package logger provides utilities for logging info
package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type slogLogger struct {
	logFile *os.File
	logger  *slog.Logger
}

// NewSlog creates slog logger
func NewSlog(logPath, mode string) (*slogLogger, error) {
	var logger *slog.Logger
	var logFile *os.File

	if mode == "dev" {
		handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
		logger = slog.New(handler)
		logFile = nil
	} else {
		logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			return nil, fmt.Errorf("failed to create logs file: %w", err)
		}
		handler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{
			Level: slog.LevelWarn,
		})
		logger = slog.New(handler)
	}

	return &slogLogger{
		logFile: logFile,
		logger:  logger,
	}, nil
}

// Error logs a message at Error level
func (s *slogLogger) Error(msg string, args ...any) {
	s.logger.Error(msg, args...)
}

// Info logs a message at Info level
func (s *slogLogger) Info(msg string, args ...any) {
	s.logger.Info(msg, args...)
}

// Debug logs a message at Debug level
func (s *slogLogger) Debug(msg string, args ...any) {
	s.logger.Debug(msg, args...)
}

// Warn logs a message at Warn level
func (s *slogLogger) Warn(msg string, args ...any) {
	s.logger.Warn(msg, args...)
}

// With returns new logger with extra key-value pair
func (s *slogLogger) With(args ...any) Logger {
	return &slogLogger{
		logFile: s.logFile,
		logger:  s.logger.With(args...),
	}
}

func (s *slogLogger) Close() error {
	if s.logFile != nil {
		err := s.logFile.Close()
		if err != nil {
			return err
		}

		s.logFile = nil
	}

	return nil
}
