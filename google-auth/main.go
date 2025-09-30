// Package main demonstrates Google OAuth integration using goth and PlainKit HTML components.
package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"google-auth/css"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"

	"github.com/plainkit/fonts/inter"
	_ "github.com/plainkit/fonts/inter/basic"
	. "github.com/plainkit/html"

	"google-auth/icons"
	"google-auth/ui/avatar"
	"google-auth/ui/button"
	"google-auth/ui/card"
	"google-auth/ui/separator"
)

// sessionUser represents a user stored in the session.
type sessionUser struct {
	Name      string
	Email     string
	AvatarURL string
}

type userContextKey struct{}

const (
	sessionName       = "plainkit-goth-session"
	providerName      = "google"
	defaultBaseURL    = "http://localhost:3000"
	defaultPort       = "3000"
	defaultSessSecret = "plainkit-goth-example-secret"
)

func main() {
	gob.Register(sessionUser{})

	clientID := mustEnv("GOOGLE_CLIENT_ID")
	clientSecret := mustEnv("GOOGLE_CLIENT_SECRET")
	baseURL := getenvDefault("BASE_URL", defaultBaseURL)
	callbackURL := baseURL + "/auth/google/callback"

	goth.UseProviders(google.New(clientID, clientSecret, callbackURL, "email", "profile"))

	sessionSecret := getenvDefault("SESSION_SECRET", defaultSessSecret)
	store := sessions.NewCookieStore([]byte(sessionSecret))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	gothic.Store = store

	mux := http.NewServeMux()
	mux.HandleFunc("/", renderHome)
	mux.HandleFunc("/auth/google", beginGoogleAuth)
	mux.HandleFunc("/auth/google/callback", completeGoogleAuth)
	mux.HandleFunc("/dashboard", requireAuth(renderDashboard))
	mux.HandleFunc("/logout", handleLogout)
	mux.HandleFunc("/assets/styles.css", cssHandler)
	inter.RegisterStatic(mux, "/assets/fonts/")

	addr := ":" + getenvDefault("PORT", defaultPort)
	fmt.Println("🚀 Goth Google OAuth Demo Server starting on " + addr)
	fmt.Println("🔗 Open http://localhost" + addr + " to view the demo")

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

// renderHome handles the home page, redirecting authenticated users to dashboard.
func renderHome(w http.ResponseWriter, r *http.Request) {
	user := currentUser(r)
	if user != nil {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	content := card.Card(
		card.Header(
			card.Title(T("Welcome to PlainKit + Goth")),
			card.Description(T("Authenticate with Google to access your personalized dashboard.")),
		),
		card.Content(
			button.Button(
				button.Props{Variant: button.VariantDefault, Size: button.SizeLg, FullWidth: true, Href: "/auth/google"},
				icons.Google(AClass("w-5 h-5")),
				T("Continue with Google"),
			),
		),
	)

	renderPage(w, "Home", content)
}

// renderDashboard displays the user dashboard with profile information.
func renderDashboard(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var avatarContent Node
	if user.AvatarURL != "" {
		avatarContent = card.Content(
			Div(
				AClass("flex justify-center mb-6"),
				avatar.Avatar(
					avatar.Props{Size: avatar.SizeLg},
					avatar.Image(avatar.ImageProps{Src: user.AvatarURL, Alt: "Profile picture"}),
					avatar.Fallback(T(string([]rune(user.Name)[0:1]))),
				),
			),
		)
	}

	content := card.Card(
		card.Header(
			card.Title(T("Welcome back, "+user.Name)),
			card.Description(T("Here's your profile information from Google.")),
		),
		avatarContent,
		card.Content(
			Div(
				AClass("space-y-4"),
				Div(
					AClass("flex justify-between items-center py-2"),
					Span(AClass("text-sm font-medium text-muted-foreground"), T("Full Name")),
					Span(AClass("text-sm font-semibold"), T(user.Name)),
				),
				separator.Separator(),
				Div(
					AClass("flex justify-between items-center py-2"),
					Span(AClass("text-sm font-medium text-muted-foreground"), T("Email Address"),
						Span(AClass("text-sm font-semibold"), T(user.Email)),
					),
				),
			),
			card.Footer(
				button.Button(
					button.Props{Variant: button.VariantOutline, Size: button.SizeDefault, FullWidth: true, Href: "/logout"},
					T("Sign Out"),
				),
			)),
	)

	renderPage(w, "Dashboard", content)
}

// renderPage renders a complete HTML page with the given title and body content.
func renderPage(w http.ResponseWriter, title string, body Node) {
	page := layout(title, body)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if _, err := w.Write([]byte("<!DOCTYPE html>\n" + Render(page))); err != nil {
		log.Printf("render error: %v", err)
	}
}

// cssHandler serves the CSS styles for the application.
func cssHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	w.Header().Set("Cache-Control", "public, max-age=31536000")

	if _, err := w.Write([]byte(css.TailwindCSS)); err != nil {
		log.Printf("write css: %v", err)
	}
}

// layout creates the base HTML structure for all pages.
func layout(title string, body Node) Node {
	return Html(
		ALang("en"),
		Head(
			Title(T(title)),
			Meta(ACharset("UTF-8")),
			Meta(AName("viewport"), AContent("width=device-width, initial-scale=1")),
			Link(ARel("stylesheet"), AHref("/assets/styles.css")),
			inter.HeadComponents("/assets/fonts"),
		),
		Body(
			AClass("min-h-screen bg-background flex items-center justify-center p-4"),
			Div(
				AClass("w-full max-w-md"),
				body,
			),
		),
	)
}

// beginGoogleAuth initiates the Google OAuth flow.
func beginGoogleAuth(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, withProvider(r, providerName))
}

// completeGoogleAuth handles the OAuth callback and creates a user session.
func completeGoogleAuth(w http.ResponseWriter, r *http.Request) {
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

	session.Values["user"] = sessionUser{
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

// handleLogout clears the user session and redirects to home.
func handleLogout(w http.ResponseWriter, r *http.Request) {
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

// requireAuth is middleware that ensures the user is authenticated.
func requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := currentUser(r)
		if user == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userContextKey{}, user)))
	}
}

// currentUser retrieves the authenticated user from the session.
func currentUser(r *http.Request) *sessionUser {
	session, err := gothic.Store.Get(r, sessionName)
	if err != nil {
		return nil
	}

	if value, ok := session.Values["user"]; ok {
		switch v := value.(type) {
		case sessionUser:
			if v.Email == "" {
				return nil
			}

			user := v

			return &user
		case *sessionUser:
			if v != nil && v.Email != "" {
				return v
			}
		}
	}

	return nil
}

// userFromContext retrieves the authenticated user from the request context.
func userFromContext(r *http.Request) *sessionUser {
	if value := r.Context().Value(userContextKey{}); value != nil {
		switch v := value.(type) {
		case *sessionUser:
			return v
		case sessionUser:
			user := v
			return &user
		}
	}

	return nil
}

// withProvider adds the provider name to the request context for goth.
func withProvider(r *http.Request, name string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), gothic.ProviderParamKey, name))
}

// mustEnv returns an environment variable or panics if it's not set.
func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("missing required environment variable %s", key)
	}

	return value
}

// getenvDefault returns an environment variable or a default value if not set.
func getenvDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
