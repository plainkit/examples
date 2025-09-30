package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/plainkit/examples/google-auth/internal/domain"
)

const (
	sessionName  = "plainkit-goth-session"
	providerName = "google"
)

// BeginAuth initiates the Google OAuth flow.
func (h *Handlers) BeginAuth(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, withProvider(r, providerName))
}

// CompleteAuth handles the OAuth callback and creates a user session.
func (h *Handlers) CompleteAuth(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, withProvider(r, providerName))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	session, err := gothic.Store.Get(r, sessionName)
	if err != nil {
		http.Error(w, "failed to load session", http.StatusInternalServerError)
		return
	}

	session.Values["user"] = domain.User{
		Name:      user.Name,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
	}
	if err := session.Save(r, w); err != nil {
		http.Error(w, "failed to save session", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// Logout clears the user session and redirects to home.
func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	if err := gothic.Logout(w, withProvider(r, providerName)); err != nil {
		log.Printf("logout error: %v", err)
	}

	session, err := gothic.Store.Get(r, sessionName)
	if err == nil {
		session.Options.MaxAge = -1
		_ = session.Save(r, w)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// withProvider adds the provider name to the request context for goth.
func withProvider(r *http.Request, name string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), gothic.ProviderParamKey, name))
}
