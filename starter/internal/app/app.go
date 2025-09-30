// Package app initializes and wires up all application components.
// This is where dependency injection happens.
package app

import (
	"log"
	"net/http"

	"starter/internal/css"
	"starter/internal/handlers"
	"starter/internal/middleware"
	"starter/internal/service"
	"starter/internal/store"
)

// App represents the application with all its dependencies.
type App struct {
	todoStore   store.TodoStore
	todoService *service.TodoService
	handlers    *handlers.Handlers
}

// New creates and initializes a new application instance.
// Dependencies are wired up from bottom-up (store -> service -> handlers).
func New() *App {
	// Initialize data store
	todoStore := store.NewMemoryStore()

	// Initialize services
	todoService := service.NewTodoService(todoStore)

	// Initialize handlers
	h := handlers.New(todoService)

	return &App{
		todoStore:   todoStore,
		todoService: todoService,
		handlers:    h,
	}
}

// Handler returns the application's HTTP handler with all routes configured.
func (a *App) Handler() http.Handler {
	mux := http.NewServeMux()

	// Static assets
	mux.HandleFunc("GET /assets/styles.css", a.serveCSS)

	// Page routes
	mux.HandleFunc("GET /{$}", a.handlers.Home)
	mux.HandleFunc("GET /todos", a.handlers.ListTodos)
	mux.HandleFunc("POST /todos", a.handlers.CreateTodo)
	mux.HandleFunc("POST /todos/{id}/complete", a.handlers.CompleteTodo)
	mux.HandleFunc("POST /todos/{id}/delete", a.handlers.DeleteTodo)

	// API routes
	mux.HandleFunc("GET /health", a.handlers.Health)

	// Apply middleware (outermost first)
	return middleware.Logger(
		middleware.Recover(mux),
	)
}

// serveCSS serves the embedded CSS stylesheet.
func (a *App) serveCSS(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=31536000")

	if _, err := w.Write([]byte(css.TailwindCSS)); err != nil {
		log.Printf("write css: %v", err)
	}
}
