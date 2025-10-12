package handlers

import (
	"net/http"

	"github.com/plainkit/bloxui/examples/starter/internal/service"
	"github.com/plainkit/bloxui/examples/starter/internal/views"
)

type Handlers struct {
	todoService *service.Todo
}

func New(todoService *service.Todo) *Handlers {
	return &Handlers{todoService: todoService}
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.List(r.Context())
	if err != nil {
		http.Error(w, "Failed to load todos", http.StatusInternalServerError)
		return
	}

	page := views.HomePage(todos)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(page))
}

func (h *Handlers) CreateTodo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form submission", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	if _, err := h.todoService.Create(r.Context(), service.CreateTodoInput{Title: title}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handlers) ToggleTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing todo id", http.StatusBadRequest)
		return
	}

	if err := h.todoService.Toggle(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handlers) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing todo id", http.StatusBadRequest)
		return
	}

	if err := h.todoService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
