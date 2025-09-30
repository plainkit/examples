package store_test

import (
	"context"
	"testing"

	"starter/internal/domain"
	"starter/internal/store"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemoryStore_Create(t *testing.T) {
	s := store.NewMemoryStore()
	ctx := context.Background()

	todo := &domain.Todo{
		ID:    "1",
		Title: "Test Todo",
	}

	err := s.Create(ctx, todo)
	require.NoError(t, err)
	assert.False(t, todo.CreatedAt.IsZero())
	assert.False(t, todo.UpdatedAt.IsZero())

	// Verify it was stored
	retrieved, err := s.Get(ctx, "1")
	require.NoError(t, err)
	assert.Equal(t, "Test Todo", retrieved.Title)
}

func TestMemoryStore_Get_NotFound(t *testing.T) {
	s := store.NewMemoryStore()
	ctx := context.Background()

	_, err := s.Get(ctx, "nonexistent")
	assert.ErrorIs(t, err, store.ErrNotFound)
}

func TestMemoryStore_List(t *testing.T) {
	s := store.NewMemoryStore()
	ctx := context.Background()

	// Create multiple todos
	todos := []*domain.Todo{
		{ID: "1", Title: "First"},
		{ID: "2", Title: "Second"},
		{ID: "3", Title: "Third"},
	}

	for _, todo := range todos {
		require.NoError(t, s.Create(ctx, todo))
	}

	// List all
	list, err := s.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 3)
}

func TestMemoryStore_Update(t *testing.T) {
	s := store.NewMemoryStore()
	ctx := context.Background()

	todo := &domain.Todo{
		ID:    "1",
		Title: "Original",
	}
	require.NoError(t, s.Create(ctx, todo))

	// Update
	todo.Title = "Updated"
	todo.Completed = true
	err := s.Update(ctx, todo)
	require.NoError(t, err)

	// Verify update
	retrieved, err := s.Get(ctx, "1")
	require.NoError(t, err)
	assert.Equal(t, "Updated", retrieved.Title)
	assert.True(t, retrieved.Completed)
}

func TestMemoryStore_Update_NotFound(t *testing.T) {
	s := store.NewMemoryStore()
	ctx := context.Background()

	todo := &domain.Todo{ID: "nonexistent"}
	err := s.Update(ctx, todo)
	assert.ErrorIs(t, err, store.ErrNotFound)
}

func TestMemoryStore_Delete(t *testing.T) {
	s := store.NewMemoryStore()
	ctx := context.Background()

	todo := &domain.Todo{ID: "1", Title: "To Delete"}
	require.NoError(t, s.Create(ctx, todo))

	// Delete
	err := s.Delete(ctx, "1")
	require.NoError(t, err)

	// Verify deletion
	_, err = s.Get(ctx, "1")
	assert.ErrorIs(t, err, store.ErrNotFound)
}

func TestMemoryStore_Delete_NotFound(t *testing.T) {
	s := store.NewMemoryStore()
	ctx := context.Background()

	err := s.Delete(ctx, "nonexistent")
	assert.ErrorIs(t, err, store.ErrNotFound)
}
