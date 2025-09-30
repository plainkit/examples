// Package main is the entry point for the application.
// It bootstraps the app and starts the HTTP server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"starter/internal/app"
)

func main() {
	// Initialize the application with all dependencies
	application := app.New()

	// Get server configuration from environment
	port := getEnv("PORT", "8080")
	addr := ":" + port

	fmt.Printf("🚀 Server starting on http://localhost:%s\n", port)

	// Start the HTTP server
	if err := http.ListenAndServe(addr, application.Handler()); err != nil {
		log.Fatal(err)
	}
}

// getEnv retrieves an environment variable or returns a fallback value.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
