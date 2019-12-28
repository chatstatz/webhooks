package tools

import "os"

// // Environment variables
// var (
// 	EnvWebhooksHost = GetEnv("WEBHOOKS_HOST", "127.0.0.0")
// 	EnvWebhooksPort = GetEnv("WEBHOOKS_PORT", "8080")
// )

// GetEnv ...
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
