package httpserver

import (
	"log/slog"

	"gamesense/apps/api/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(cfg config.Config, logger *slog.Logger) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(requestLogMiddleware(logger))
	router.Use(corsMiddleware(cfg.AllowedOrigins))

	router.Get("/health", handleHealth(cfg))

	return router
}
