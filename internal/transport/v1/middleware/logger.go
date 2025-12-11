// Package middle provides server middlewares
package middle

import (
	"context"
	"net/http"

	"github.com/TapokGo/tapok-drive/internal/logger"
	"github.com/google/uuid"
)

type contextKey struct{}

var loggerKey = contextKey{}

// LoggingMiddleware enriches base logger with request id
func LoggingMiddleware(baseLogger logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.New().String()
			logger := baseLogger.With("requestID", requestID)
			ctx := context.WithValue(r.Context(), loggerKey, logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// FromContext receives logger from context safely
func FromContext(ctx context.Context) logger.Logger {
	if logger, ok := ctx.Value(loggerKey).(logger.Logger); ok {
		return logger
	}

	// Use panic cuz its programmer mistake
	panic("logger not found in context")
}
