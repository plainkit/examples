package handlers

import (
	"net/http"

	"starter/internal/service"
	"starter/internal/views"
)

// ListTodos displays all todos.
func (h *Handlers) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoService.List(r.Context())
	if err != nil {
		http.Error(w, "Failed to load todos", http.StatusInternalServerError)
		return
	}

	h.renderPage(w, views.TodosPage(todos))
}

// CreateTodo handles todo creation.
func (h *Handlers) CreateTodo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")

	_, err := h.todoService.Create(r.Context(), service.CreateTodoInput{
		Title: title,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Redirect back to todo list
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

// CompleteTodo marks a todo as completed.
func (h *Handlers) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// Get existing todo
	todo, err := h.todoService.Get(r.Context(), id)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// Update to completed
	_, err = h.todoService.Update(r.Context(), id, service.UpdateTodoInput{
		Title:     todo.Title,
		Completed: true,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Redirect back to todo list
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

// DeleteTodo removes a todo.
func (h *Handlers) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.todoService.Delete(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}

	// Redirect back to todo list
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}
