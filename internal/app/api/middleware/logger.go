package middleware

import (
	"butta/pkg/logger"
	"net/http"
	"time"
)

// Logger is a middleware handler that logs the request method, URL, and processing time.
func Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Info("Request processed", "method", r.Method, "uri", r.RequestURI, "duration", time.Since(start))
	}
}
