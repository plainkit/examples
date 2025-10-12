package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/plainkit/bloxui/examples/starter/internal/domain"
	"github.com/plainkit/bloxui/examples/starter/internal/store"
)

var (
	ErrTitleRequired = errors.New("title is required")
)

type Todo struct {
	store store.TodoStore
}

func NewTodoService(s store.TodoStore) *Todo {
	return &Todo{store: s}
}

type CreateTodoInput struct {
	Title string
}

func (t *Todo) Create(ctx context.Context, input CreateTodoInput) (*domain.Todo, error) {
	title := strings.TrimSpace(input.Title)
	if title == "" {
		return nil, ErrTitleRequired
	}

	now := time.Now().UTC()
	todo := &domain.Todo{
		ID:        uuid.NewString(),
		Title:     title,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := t.store.Create(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *Todo) Toggle(ctx context.Context, id string) error {
	todo, err := t.store.Get(ctx, id)
	if err != nil {
		return err
	}

	todo.Completed = !todo.Completed
	todo.UpdatedAt = time.Now().UTC()

	return t.store.Update(ctx, todo)
}

func (t *Todo) Delete(ctx context.Context, id string) error {
	return t.store.Delete(ctx, id)
}

func (t *Todo) List(ctx context.Context) ([]*domain.Todo, error) {
	return t.store.List(ctx)
}
