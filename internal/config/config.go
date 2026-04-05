package config

import (
	"log" // error logging
	"os"  // environment variable handling

	"github.com/joho/godotenv" // for loading .env files into environment variables
)

// Config holds the application configuration values
type Config struct {
	AppEnv      string
	Port        string
	DatabaseURL string
	ClientURL   string
}

// getEnv retrieves the value of the environment variable named by the key.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// mustGetEnv retrieves the value of the environment variable named by the key.
// If the variable is not set, it logs a fatal error and exits the application.
func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s is required", key)
	}
	return value
}

// LoadConfig loads the configuration from environment variables and returns a Config struct.
func LoadConfig() *Config {
	_ = godotenv.Load() // Load .env file if it exists, ignore error if it doesn't

	return &Config{
		AppEnv:      getEnv("APP_ENV", "development"),
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: mustGetEnv("DATABASE_URL"),
		ClientURL:   getEnv("CLIENT_URL", "http://localhost:3000"),
	}
}
