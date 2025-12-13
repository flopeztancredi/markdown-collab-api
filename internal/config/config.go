package config

import "os"

type Config struct {
	AppName    string
	AppPort    string
	AppVersion string
	GinMode    string
}

func Load() *Config {
	return &Config{
		AppName:    getEnv("APP_NAME", "markdown-collab"),
		AppPort:    getEnv("APP_PORT", "8080"),
		AppVersion: getEnv("APP_VERSION", "1.0.0"),
		GinMode:    getEnv("GIN_MODE", "release"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
