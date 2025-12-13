package config

import "os"

type Config struct {
	AppName     string
	AppPort     string
	GinMode     string
	DatabaseURL string
}

func Load() *Config {
	return &Config{
		AppName:     getEnv("APP_NAME", "markdown-collab-api"),
		AppPort:     getEnv("APP_PORT", "8080"),
		GinMode:     getEnv("GIN_MODE", "debug"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/markdown?sslmode=disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
