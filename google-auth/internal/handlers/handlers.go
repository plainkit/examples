package handlers

import (
	"log"
	"net/http"

	"github.com/plainkit/examples/google-auth/internal/middleware"
	"github.com/plainkit/examples/google-auth/internal/views"
)

// Handlers contains all HTTP handlers.
type Handlers struct{}

// New creates a new Handlers instance.
func New() *Handlers {
	return &Handlers{}
}

// Home handles the home page, redirecting authenticated users to dashboard.
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	user := middleware.CurrentUser(r)
	if user != nil {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	h.renderPage(w, views.HomePage())
}

// Dashboard displays the user dashboard with profile information.
func (h *Handlers) Dashboard(w http.ResponseWriter, r *http.Request) {
	user := middleware.UserFromContext(r)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	h.renderPage(w, views.DashboardPage(user))
}

// renderPage renders a complete HTML page.
func (h *Handlers) renderPage(w http.ResponseWriter, content string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if _, err := w.Write([]byte(content)); err != nil {
		log.Printf("render error: %v", err)
	}
}
