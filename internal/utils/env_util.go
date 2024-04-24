package utils

import (
	"os"

	"github.com/lessapidev/lessapi-duckduckgo/internal/envs"
)

// GetEnvOrDefault returns the value of the environment variable with the given key.
func GetEnvOrDefault(key envs.EnvKey, defaultValue string) string {
	value := os.Getenv(string(key))
	if value == "" {
		return defaultValue
	}
	return value
}

// GetEnv returns the value of the environment variable with the given key.
func GetEnv(key envs.EnvKey) string {
	return os.Getenv(string(key))
}
