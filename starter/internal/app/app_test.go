package app_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/plainkit/bloxui/examples/starter/internal/app"
)

func TestTodoFlow(t *testing.T) {
	application := app.New()
	server := httptest.NewServer(application.Handler())
	t.Cleanup(server.Close)

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	form := url.Values{}
	form.Set("title", "Design landing page")

	resp, err := client.Post(server.URL+"/todos", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	require.NoError(t, err)
	require.Equal(t, http.StatusSeeOther, resp.StatusCode)
	_ = resp.Body.Close()

	resp, err = http.Get(server.URL + "/")
	require.NoError(t, err)

	defer func() { _ = resp.Body.Close() }()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Contains(t, string(body), "Design landing page")
}
