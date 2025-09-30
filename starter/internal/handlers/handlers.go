// Package handlers contains HTTP request handlers.
// Handlers parse requests, call services, and render responses.
package handlers

import (
	"log"
	"net/http"

	"starter/internal/service"
	"starter/internal/views"
)

// Handlers contains all HTTP handlers and their dependencies.
type Handlers struct {
	todoService *service.TodoService
}

// New creates a new Handlers instance with injected dependencies.
func New(todoService *service.TodoService) *Handlers {
	return &Handlers{
		todoService: todoService,
	}
}

// Home renders the home page.
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	h.renderPage(w, views.HomePage())
}

// Health returns a simple health check response.
func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

// renderPage renders a complete HTML page.
func (h *Handlers) renderPage(w http.ResponseWriter, content string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if _, err := w.Write([]byte(content)); err != nil {
		log.Printf("render error: %v", err)
	}
}
