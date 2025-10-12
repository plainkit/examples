package store

import (
	"context"
	"errors"
	"sort"
	"sync"

	"github.com/plainkit/bloxui/examples/starter/internal/domain"
)

type TodoStore interface {
	Create(ctx context.Context, todo *domain.Todo) error
	Get(ctx context.Context, id string) (*domain.Todo, error)
	List(ctx context.Context) ([]*domain.Todo, error)
	Update(ctx context.Context, todo *domain.Todo) error
	Delete(ctx context.Context, id string) error
}

var ErrNotFound = errors.New("todo not found")

type MemoryStore struct {
	mu    sync.RWMutex
	todos map[string]*domain.Todo
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{todos: make(map[string]*domain.Todo)}
}

func (s *MemoryStore) Create(_ context.Context, todo *domain.Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.todos[todo.ID] = todo

	return nil
}

func (s *MemoryStore) Get(_ context.Context, id string) (*domain.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, ok := s.todos[id]
	if !ok {
		return nil, ErrNotFound
	}

	clone := *todo

	return &clone, nil
}

func (s *MemoryStore) List(_ context.Context) ([]*domain.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todos := make([]*domain.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		clone := *todo
		todos = append(todos, &clone)
	}

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].CreatedAt.After(todos[j].CreatedAt)
	})

	return todos, nil
}

func (s *MemoryStore) Update(_ context.Context, todo *domain.Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.todos[todo.ID]; !ok {
		return ErrNotFound
	}

	s.todos[todo.ID] = todo

	return nil
}

func (s *MemoryStore) Delete(_ context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.todos[id]; !ok {
		return ErrNotFound
	}

	delete(s.todos, id)

	return nil
}
