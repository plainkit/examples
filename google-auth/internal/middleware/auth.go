package middleware

import (
	"context"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/plainkit/examples/google-auth/internal/domain"
)

type userContextKey struct{}

const sessionName = "plainkit-goth-session"

// RequireAuth is middleware that ensures the user is authenticated.
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := CurrentUser(r)
		if user == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userContextKey{}, user)))
	}
}

// CurrentUser retrieves the authenticated user from the session.
func CurrentUser(r *http.Request) *domain.User {
	session, err := gothic.Store.Get(r, sessionName)
	if err != nil {
		return nil
	}

	if value, ok := session.Values["user"]; ok {
		switch v := value.(type) {
		case domain.User:
			if v.Email == "" {
				return nil
			}

			user := v

			return &user
		case *domain.User:
			if v != nil && v.Email != "" {
				return v
			}
		}
	}

	return nil
}

// UserFromContext retrieves the authenticated user from the request context.
func UserFromContext(r *http.Request) *domain.User {
	if value := r.Context().Value(userContextKey{}); value != nil {
		switch v := value.(type) {
		case *domain.User:
			return v
		case domain.User:
			user := v
			return &user
		}
	}

	return nil
}
