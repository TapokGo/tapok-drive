// Package middle provides server middlewares
package middle

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/TapokGo/tapok-drive/internal/logger"
	"github.com/google/uuid"
)

type contextKey struct{}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *statusRecorder) WriteHeader(statusCode int) {
	rec.statusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

var loggerKey = contextKey{}

// LoggingMiddleware enriches base logger with request id
func RequestLogger(baseLogger logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/healthz" || strings.HasPrefix(r.URL.Path, "/swagger/") {
				next.ServeHTTP(w, r)
				return
			}

			// Create enriches logger
			requestID := uuid.New().String()
			logger := baseLogger.With("requestID", requestID)
			ctx := context.WithValue(r.Context(), loggerKey, logger)

			// Add recorder
			recorder := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}
			start := time.Now()
			next.ServeHTTP(recorder, r.WithContext(ctx))

			// Logging exit
			duration := time.Since(start)

			logger.Info("request complete", "method", r.Method, "path", r.URL.Path, "status", recorder.statusCode, "duration", duration)
			logger.Debug("------------------------------------------------------------------")
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
