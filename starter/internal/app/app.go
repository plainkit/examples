package app

import (
	"net/http"

	"github.com/plainkit/bloxui/examples/starter/internal/css"
	"github.com/plainkit/bloxui/examples/starter/internal/handlers"
	"github.com/plainkit/bloxui/examples/starter/internal/middleware"
	"github.com/plainkit/bloxui/examples/starter/internal/service"
	"github.com/plainkit/bloxui/examples/starter/internal/store"
)

type App struct {
	todoStore   store.TodoStore
	todoService *service.Todo
	handlers    *handlers.Handlers
}

func New() *App {
	todoStore := store.NewMemoryStore()
	todoService := service.NewTodoService(todoStore)
	handlers := handlers.New(todoService)

	return &App{
		todoStore:   todoStore,
		todoService: todoService,
		handlers:    handlers,
	}
}

func (a *App) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /assets/styles.css", serveCSS)

	mux.HandleFunc("GET /{$}", a.handlers.Home)
	mux.HandleFunc("POST /todos", a.handlers.CreateTodo)
	mux.HandleFunc("POST /todos/{id}/toggle", a.handlers.ToggleTodo)
	mux.HandleFunc("POST /todos/{id}/delete", a.handlers.DeleteTodo)

	return middleware.Logger(middleware.Recover(mux))
}

func serveCSS(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=3600")
	_, _ = w.Write([]byte(css.TailwindCSS))
}
