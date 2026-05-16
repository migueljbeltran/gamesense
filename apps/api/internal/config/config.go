package config

import "os"

type Config struct {
	Environment    string
	LogLevel       string
	APIAddr        string
	AllowedOrigins string
	DatabaseURL    string
}

func Load() Config {
	return Config{
		Environment:    getEnv("ENVIRONMENT", "development"),
		LogLevel:       getEnv("LOG_LEVEL", "debug"),
		APIAddr:        getEnv("API_ADDR", ":8000"),
		AllowedOrigins: getEnv("API_CORS_ORIGINS", "http://localhost:3000"),
		DatabaseURL:    getEnv("DATABASE_URL", "postgres://gamesense:gamesense@localhost:5432/gamesense?sslmode=disable"),
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
