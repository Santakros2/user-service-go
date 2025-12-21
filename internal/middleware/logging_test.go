package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"users-service/internal/middleware"
)

func TestLoggingMiddleware_CallsNextHandler(t *testing.T) {

	// flag to see if handler was called
	called := false

	// dummy handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	// wrap with middleware
	handler := middleware.Logging(nextHandler)

	// fake request
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()

	// execute
	handler.ServeHTTP(rec, req)

	// assertions
	if !called {
		t.Fatal("expected next handler to be called")
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
}
