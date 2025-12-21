package middleware

import (
	"net/http"
	"time"
	"users-service/internal/logger"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logger.Logger.Printf(
			"START %s %s",
			r.Method,
			r.URL.Path,
		)

		next.ServeHTTP(w, r)

		logger.Logger.Printf(
			"END   %s %s (%v)",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
