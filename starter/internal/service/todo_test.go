package service_test

import (
	"context"
	"testing"

	"starter/internal/service"
	"starter/internal/store"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTodoService_Create(t *testing.T) {
	todoStore := store.NewMemoryStore()
	svc := service.NewTodoService(todoStore)
	ctx := context.Background()

	t.Run("valid todo", func(t *testing.T) {
		todo, err := svc.Create(ctx, service.CreateTodoInput{
			Title: "Buy groceries",
		})

		require.NoError(t, err)
		assert.NotEmpty(t, todo.ID)
		assert.Equal(t, "Buy groceries", todo.Title)
		assert.False(t, todo.Completed)
	})

	t.Run("empty title", func(t *testing.T) {
		_, err := svc.Create(ctx, service.CreateTodoInput{
			Title: "",
		})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "title is required")
	})

	t.Run("whitespace title", func(t *testing.T) {
		_, err := svc.Create(ctx, service.CreateTodoInput{
			Title: "   ",
		})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "title is required")
	})
}

func TestTodoService_List(t *testing.T) {
	todoStore := store.NewMemoryStore()
	svc := service.NewTodoService(todoStore)
	ctx := context.Background()

	// Create multiple todos
	_, err := svc.Create(ctx, service.CreateTodoInput{Title: "First"})
	require.NoError(t, err)
	_, err = svc.Create(ctx, service.CreateTodoInput{Title: "Second"})
	require.NoError(t, err)

	// List all
	todos, err := svc.List(ctx)
	require.NoError(t, err)
	assert.Len(t, todos, 2)
}

func TestTodoService_Get(t *testing.T) {
	todoStore := store.NewMemoryStore()
	svc := service.NewTodoService(todoStore)
	ctx := context.Background()

	// Create a todo
	created, err := svc.Create(ctx, service.CreateTodoInput{Title: "Test"})
	require.NoError(t, err)

	// Get it back
	todo, err := svc.Get(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, todo.ID)
	assert.Equal(t, "Test", todo.Title)
}

func TestTodoService_Update(t *testing.T) {
	todoStore := store.NewMemoryStore()
	svc := service.NewTodoService(todoStore)
	ctx := context.Background()

	// Create a todo
	created, err := svc.Create(ctx, service.CreateTodoInput{Title: "Original"})
	require.NoError(t, err)

	t.Run("valid update", func(t *testing.T) {
		updated, err := svc.Update(ctx, created.ID, service.UpdateTodoInput{
			Title:     "Updated",
			Completed: true,
		})

		require.NoError(t, err)
		assert.Equal(t, "Updated", updated.Title)
		assert.True(t, updated.Completed)
	})

	t.Run("empty title", func(t *testing.T) {
		_, err := svc.Update(ctx, created.ID, service.UpdateTodoInput{
			Title: "",
		})

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "title is required")
	})

	t.Run("not found", func(t *testing.T) {
		_, err := svc.Update(ctx, "nonexistent", service.UpdateTodoInput{
			Title: "Test",
		})

		assert.ErrorIs(t, err, store.ErrNotFound)
	})
}

func TestTodoService_Delete(t *testing.T) {
	todoStore := store.NewMemoryStore()
	svc := service.NewTodoService(todoStore)
	ctx := context.Background()

	// Create a todo
	created, err := svc.Create(ctx, service.CreateTodoInput{Title: "To Delete"})
	require.NoError(t, err)

	// Delete it
	err = svc.Delete(ctx, created.ID)
	require.NoError(t, err)

	// Verify it's gone
	_, err = svc.Get(ctx, created.ID)
	assert.ErrorIs(t, err, store.ErrNotFound)
}
