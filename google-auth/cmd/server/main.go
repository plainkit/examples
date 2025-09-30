// Package main demonstrates Google OAuth integration using goth and PlainKit HTML components.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/plainkit/examples/google-auth/internal/app"
)

func main() {
	application := app.New()

	port := getEnv("PORT", "3000")
	addr := ":" + port

	fmt.Printf("🚀 Goth Google OAuth Demo Server starting on %s\n", addr)
	fmt.Printf("🔗 Open http://localhost:%s to view the demo\n", port)

	if err := http.ListenAndServe(addr, application.Handler()); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
