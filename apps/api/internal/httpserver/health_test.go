package httpserver

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"gamesense/apps/api/internal/config"
)

func TestHealth(t *testing.T) {
	cfg := config.Config{
		Environment:    "test",
		AllowedOrigins: "http://localhost:3000",
	}
	router := NewRouter(cfg, slog.Default())

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	var payload healthResponse
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if payload.Status != "ok" {
		t.Fatalf("expected status ok, got %q", payload.Status)
	}

	if payload.Environment != "test" {
		t.Fatalf("expected environment test, got %q", payload.Environment)
	}
}
