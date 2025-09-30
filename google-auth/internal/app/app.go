package app

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/plainkit/fonts/inter"

	"github.com/plainkit/examples/google-auth/internal/css"
	"github.com/plainkit/examples/google-auth/internal/domain"
	"github.com/plainkit/examples/google-auth/internal/handlers"
	"github.com/plainkit/examples/google-auth/internal/middleware"
)

const (
	defaultBaseURL    = "http://localhost:3000"
	defaultSessSecret = "plainkit-goth-example-secret"
)

// App represents the application with all its dependencies.
type App struct {
	handlers *handlers.Handlers
}

// New creates and initializes a new application instance.
func New() *App {
	// Register domain.User for session encoding
	gob.Register(domain.User{})

	// Configure Google OAuth provider
	clientID := mustEnv("GOOGLE_CLIENT_ID")
	clientSecret := mustEnv("GOOGLE_CLIENT_SECRET")
	baseURL := getEnvDefault("BASE_URL", defaultBaseURL)
	callbackURL := baseURL + "/auth/google/callback"

	goth.UseProviders(google.New(clientID, clientSecret, callbackURL, "email", "profile"))

	// Configure session store
	sessionSecret := getEnvDefault("SESSION_SECRET", defaultSessSecret)
	store := sessions.NewCookieStore([]byte(sessionSecret))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	gothic.Store = store

	// Initialize handlers
	h := handlers.New()

	return &App{
		handlers: h,
	}
}

// Handler returns the application's HTTP handler with all routes configured.
func (a *App) Handler() http.Handler {
	mux := http.NewServeMux()

	// Static assets
	mux.HandleFunc("GET /assets/styles.css", a.serveCSS)
	inter.RegisterStatic(mux, "/assets/fonts/")

	// Auth routes
	mux.HandleFunc("GET /auth/google", a.handlers.BeginAuth)
	mux.HandleFunc("GET /auth/google/callback", a.handlers.CompleteAuth)
	mux.HandleFunc("GET /logout", a.handlers.Logout)

	// Page routes
	mux.HandleFunc("GET /{$}", a.handlers.Home)
	mux.HandleFunc("GET /dashboard", middleware.RequireAuth(a.handlers.Dashboard))

	return mux
}

// serveCSS serves the embedded CSS stylesheet.
func (a *App) serveCSS(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=31536000")

	if _, err := w.Write([]byte(css.TailwindCSS)); err != nil {
		log.Printf("write css: %v", err)
	}
}

// mustEnv returns an environment variable or panics if it's not set.
func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("missing required environment variable %s", key)
	}

	return value
}

// getEnvDefault returns an environment variable or a default value if not set.
func getEnvDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
