// Package service contains business logic and orchestration.
// Services validate input, enforce business rules, and coordinate between stores.
package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"starter/internal/domain"
	"starter/internal/store"
)

// TodoService handles todo business logic.
type TodoService struct {
	store store.TodoStore
}

// NewTodoService creates a new todo service.
func NewTodoService(store store.TodoStore) *TodoService {
	return &TodoService{store: store}
}

// CreateTodoInput represents input for creating a todo.
type CreateTodoInput struct {
	Title string
}

// Create validates input and creates a new todo.
func (s *TodoService) Create(ctx context.Context, input CreateTodoInput) (*domain.Todo, error) {
	// Validation
	if strings.TrimSpace(input.Title) == "" {
		return nil, errors.New("title is required")
	}

	// Business logic: generate ID
	id := generateID()

	// Create todo entity
	todo := &domain.Todo{
		ID:        id,
		Title:     input.Title,
		Completed: false,
	}

	// Persist
	if err := s.store.Create(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

// List retrieves all todos.
func (s *TodoService) List(ctx context.Context) ([]*domain.Todo, error) {
	return s.store.List(ctx)
}

// Get retrieves a todo by ID.
func (s *TodoService) Get(ctx context.Context, id string) (*domain.Todo, error) {
	return s.store.Get(ctx, id)
}

// UpdateTodoInput represents input for updating a todo.
type UpdateTodoInput struct {
	Title     string
	Completed bool
}

// Update validates input and updates an existing todo.
func (s *TodoService) Update(ctx context.Context, id string, input UpdateTodoInput) (*domain.Todo, error) {
	// Validation
	if strings.TrimSpace(input.Title) == "" {
		return nil, errors.New("title is required")
	}

	// Get existing todo
	todo, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields
	todo.Title = input.Title
	todo.Completed = input.Completed

	// Persist
	if err := s.store.Update(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

// Delete removes a todo by ID.
func (s *TodoService) Delete(ctx context.Context, id string) error {
	return s.store.Delete(ctx, id)
}

// generateID creates a unique ID for a todo.
// In production, use a proper ID generator (UUID, ULID, etc.)
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
