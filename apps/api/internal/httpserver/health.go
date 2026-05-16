package httpserver

import (
	"encoding/json"
	"net/http"

	"gamesense/apps/api/internal/config"
)

type healthResponse struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
}

func handleHealth(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, healthResponse{
			Status:      "ok",
			Environment: cfg.Environment,
		})
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
