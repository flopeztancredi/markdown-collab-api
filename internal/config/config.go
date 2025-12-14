package config

import (
	"os"
	"strings"
	"time"
)

type Config struct {
	AppName        string
	AppPort        string
	GinMode        string
	DatabaseURL    string
	AllowedOrigins []string
	RequestTimeout time.Duration
	DBTimeout      time.Duration
}

func Load() *Config {
	return &Config{
		AppName:        getEnv("APP_NAME", "markdown-collab-api"),
		AppPort:        getEnv("APP_PORT", "8080"),
		GinMode:        getEnv("GIN_MODE", "debug"),
		DatabaseURL:    getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/markdown?sslmode=disable"),
		AllowedOrigins: splitAndTrim(getEnv("ALLOWED_ORIGINS", "http://localhost:3000,http://localhost:5173")),
		RequestTimeout: getEnvDuration("REQUEST_TIMEOUT", "15s"),
		DBTimeout:      getEnvDuration("DB_TIMEOUT", "5s"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func splitAndTrim(value string) []string {
	raw := strings.Split(value, ",")
	var out []string
	for _, v := range raw {
		trimmed := strings.TrimSpace(v)
		if trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}

func getEnvDuration(key, defaultValue string) time.Duration {
	raw := getEnv(key, defaultValue)
	d, err := time.ParseDuration(raw)
	if err != nil {
		d, _ = time.ParseDuration(defaultValue)
	}
	return d
}
