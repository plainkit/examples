package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/plainkit/bloxui/examples/starter/internal/app"
)

func main() {
	application := app.New()

	port := getenv("PORT", "8080")
	addr := ":" + port

	fmt.Printf("Server listening on http://localhost:%s\n", port)

	if err := http.ListenAndServe(addr, application.Handler()); err != nil {
		log.Fatal(err)
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
