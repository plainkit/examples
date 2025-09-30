// Package app_test demonstrates integration testing.
// Integration tests are preferred over heavy unit testing.
package app_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"starter/internal/app"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestHomePage tests the home page renders correctly.
func TestHomePage(t *testing.T) {
	application := app.New()

	server := httptest.NewServer(application.Handler())
	defer server.Close()

	resp, err := http.Get(server.URL + "/")
	require.NoError(t, err)

	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	assert.Contains(t, string(body), "Welcome to PlainKit Starter")
}

// TestHealthCheck tests the health endpoint.
func TestHealthCheck(t *testing.T) {
	application := app.New()

	server := httptest.NewServer(application.Handler())
	defer server.Close()

	resp, err := http.Get(server.URL + "/health")
	require.NoError(t, err)

	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	assert.JSONEq(t, `{"status":"ok"}`, string(body))
}

// TestTodoFlow tests the complete todo workflow.
func TestTodoFlow(t *testing.T) {
	application := app.New()

	server := httptest.NewServer(application.Handler())
	defer server.Close()

	// 1. List todos (should be empty)
	resp, err := http.Get(server.URL + "/todos")
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	assert.Contains(t, string(body), "No todos yet")

	_ = resp.Body.Close()

	// 2. Create a todo
	resp, err = http.Post(
		server.URL+"/todos",
		"application/x-www-form-urlencoded",
		strings.NewReader("title=Buy+groceries"),
	)
	require.NoError(t, err)
	// http.Post follows redirects, so we get the final 200, not the 303
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	_ = resp.Body.Close()

	// 3. Verify todo appears in the list
	resp, err = http.Get(server.URL + "/todos")
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, _ = io.ReadAll(resp.Body)
	assert.Contains(t, string(body), "Buy groceries")
	assert.NotContains(t, string(body), "No todos yet")

	_ = resp.Body.Close()
}

// TestCreateTodo_Validation tests todo creation validation.
func TestCreateTodo_Validation(t *testing.T) {
	application := app.New()

	server := httptest.NewServer(application.Handler())
	defer server.Close()

	// Try to create todo with empty title
	resp, err := http.Post(
		server.URL+"/todos",
		"application/x-www-form-urlencoded",
		strings.NewReader("title="),
	)
	require.NoError(t, err)

	defer func() { _ = resp.Body.Close() }()

	// Should get a bad request
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	assert.Contains(t, string(body), "title is required")
}

// TestStaticAssets tests CSS is served correctly.
func TestStaticAssets(t *testing.T) {
	application := app.New()

	server := httptest.NewServer(application.Handler())
	defer server.Close()

	resp, err := http.Get(server.URL + "/assets/styles.css")
	require.NoError(t, err)

	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "text/css; charset=utf-8", resp.Header.Get("Content-Type"))
}
