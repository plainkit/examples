package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/plainkit/bloxui/examples/starter/internal/service"
	"github.com/plainkit/bloxui/examples/starter/internal/store"
)

func TestTodoService_Create(t *testing.T) {
	svc := service.NewTodoService(store.NewMemoryStore())

	todo, err := svc.Create(context.Background(), service.CreateTodoInput{Title: "Ship starter"})
	require.NoError(t, err)
	require.Equal(t, "Ship starter", todo.Title)
	require.NotEmpty(t, todo.ID)
}

func TestTodoService_Toggle(t *testing.T) {
	svc := service.NewTodoService(store.NewMemoryStore())

	todo, err := svc.Create(context.Background(), service.CreateTodoInput{Title: "Write docs"})
	require.NoError(t, err)

	err = svc.Toggle(context.Background(), todo.ID)
	require.NoError(t, err)

	list, err := svc.List(context.Background())
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.True(t, list[0].Completed)
}

func TestTodoService_Delete(t *testing.T) {
	svc := service.NewTodoService(store.NewMemoryStore())

	todo, err := svc.Create(context.Background(), service.CreateTodoInput{Title: "Write docs"})
	require.NoError(t, err)

	err = svc.Delete(context.Background(), todo.ID)
	require.NoError(t, err)

	list, err := svc.List(context.Background())
	require.NoError(t, err)
	require.Len(t, list, 0)
}

func TestTodoService_RequiresTitle(t *testing.T) {
	svc := service.NewTodoService(store.NewMemoryStore())

	_, err := svc.Create(context.Background(), service.CreateTodoInput{Title: "   "})
	require.ErrorIs(t, err, service.ErrTitleRequired)
}
