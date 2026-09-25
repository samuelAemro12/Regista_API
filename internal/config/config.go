package config

import (
	"fmt"
	"os"
)

// Config contains runtime settings read from environment variables.
type Config struct {
	AppEnv      string
	Port        string
	DatabaseURL string
	RedisURL    string
}

// Load reads application configuration. DatabaseURL is intentionally required so
// credentials are supplied by the environment rather than embedded in the app.
func Load() (Config, error) {
	cfg := Config{
		AppEnv:      getEnv("APP_ENV", "development"),
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		RedisURL:    getEnv("REDIS_URL", "redis://localhost:6379/0"),
	}

	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
