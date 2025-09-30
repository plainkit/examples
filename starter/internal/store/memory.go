// Package store provides data persistence implementations.
// This package defines interfaces and implementations for data access.
package store

import (
	"context"
	"errors"
	"sync"
	"time"

	"starter/internal/domain"
)

var (
	// ErrNotFound is returned when a todo is not found.
	ErrNotFound = errors.New("todo not found")
)

// TodoStore defines the interface for todo data operations.
type TodoStore interface {
	Create(ctx context.Context, todo *domain.Todo) error
	Get(ctx context.Context, id string) (*domain.Todo, error)
	List(ctx context.Context) ([]*domain.Todo, error)
	Update(ctx context.Context, todo *domain.Todo) error
	Delete(ctx context.Context, id string) error
}

// MemoryStore is an in-memory implementation of TodoStore.
// Useful for demos, testing, and development.
type MemoryStore struct {
	mu    sync.RWMutex
	todos map[string]*domain.Todo
}

// NewMemoryStore creates a new in-memory todo store.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		todos: make(map[string]*domain.Todo),
	}
}

// Create adds a new todo to the store.
func (s *MemoryStore) Create(ctx context.Context, todo *domain.Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	s.todos[todo.ID] = todo

	return nil
}

// Get retrieves a todo by ID.
func (s *MemoryStore) Get(ctx context.Context, id string) (*domain.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, exists := s.todos[id]
	if !exists {
		return nil, ErrNotFound
	}

	return todo, nil
}

// List returns all todos.
func (s *MemoryStore) List(ctx context.Context) ([]*domain.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todos := make([]*domain.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}

	return todos, nil
}

// Update modifies an existing todo.
func (s *MemoryStore) Update(ctx context.Context, todo *domain.Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.todos[todo.ID]; !exists {
		return ErrNotFound
	}

	todo.UpdatedAt = time.Now()
	s.todos[todo.ID] = todo

	return nil
}

// Delete removes a todo by ID.
func (s *MemoryStore) Delete(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.todos[id]; !exists {
		return ErrNotFound
	}

	delete(s.todos, id)

	return nil
}
