// Package main demonstrates Google OAuth integration using goth and PlainKit HTML components.
package main

import (
	"context"
	"encoding/gob"
	"fmt"
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

	content := []Node{
		H1(T("PlainKit + Goth")),
		P(AClass("muted"), T("Authenticate with Google to access the private dashboard.")),
		Div(
			AClass("actions"),
			A(AHref("/auth/google"), AClass("button"),
				T("Continue with Google"),
			),
		),
	}

	renderPage(w, "Home", content...)
}

// renderDashboard displays the user dashboard with profile information.
func renderDashboard(w http.ResponseWriter, r *http.Request) {
	user := userFromContext(r)
	if user == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	content := []Node{
		H1(T("Welcome back")),
	}

	if user.AvatarURL != "" {
		content = append(content,
			Div(
				AClass("avatar"),
				Img(ASrc(user.AvatarURL), AAlt("Avatar")),
			),
		)
	}

	content = append(content,
		Div(
			AClass("surface"),
			Div(
				AClass("data-row"),
				Span(AClass("data-label"), T("Name")),
				Strong(T(user.Name)),
			),
			Div(
				AClass("data-row"),
				Span(AClass("data-label"), T("Email")),
				Strong(T(user.Email)),
			),
		),
		Div(
			AClass("actions"),
			A(AHref("/logout"), AClass("button"), T("Sign out")),
		),
	)

	renderPage(w, "Dashboard", content...)
}

// renderPage renders a complete HTML page with the given title and body content.
func renderPage(w http.ResponseWriter, title string, body ...Node) {
	page := layout(title, body...)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := w.Write([]byte("<!DOCTYPE html>\n" + Render(page))); err != nil {
		log.Printf("render error: %v", err)
	}
}

// layout creates the base HTML structure for all pages.
func layout(title string, body ...Node) Node {
	cardChildren := make([]DivArg, 0, len(body)+1)
	cardChildren = append(cardChildren, AClass("card shell"))
	for _, node := range body {
		cardChildren = append(cardChildren, Child(node))
	}

	return Html(
		ALang("en"),
		Head(
			Title(T(title)),
			Meta(ACharset("UTF-8")),
			Meta(AName("viewport"), AContent("width=device-width, initial-scale=1")),
			inter.HeadComponents("/assets/fonts"),
			Style(UnsafeText(baseStyles())),
		),
		Body(
			AClass("page"),
			Main(
				AClass("container"),
				Div(cardChildren...),
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

// baseStyles returns the CSS styles for the application.
func baseStyles() string {
	return `:root {
  color-scheme: light dark;
  --bg: #fafafa; /* zinc-50 */
  --bg-gradient: radial-gradient(circle at top left, rgba(113, 113, 122, 0.08), rgba(161, 161, 170, 0.05));
  --fg: #18181b; /* zinc-900 */
  --muted: #71717a; /* zinc-500 */
  --card-bg: #ffffff; /* white */
  --surface-bg: #fafafa; /* zinc-50 */
  --border: rgba(161, 161, 170, 0.2); /* zinc-400 with opacity */
  --shadow: 0 28px 80px rgba(24, 24, 27, 0.12); /* zinc-900 shadow */
  --surface-shadow: 0 20px 50px rgba(24, 24, 27, 0.06);
  --accent: #3f3f46; /* zinc-700 */
  --accent-hover: #27272a; /* zinc-800 */
  --button-bg: rgba(63, 63, 70, 0.08); /* zinc-700 with opacity */
  --button-border: rgba(113, 113, 122, 0.25); /* zinc-500 with opacity */
  --button-hover-bg: rgba(63, 63, 70, 0.15); /* zinc-700 with opacity */
  --avatar-border: rgba(113, 113, 122, 0.4); /* zinc-500 with opacity */
  --avatar-shadow: 0 10px 30px rgba(24, 24, 27, 0.18); /* zinc-900 shadow */
}
@media (prefers-color-scheme: dark) {
  :root {
    --bg: #09090b; /* zinc-950 */
    --bg-gradient: radial-gradient(circle at top left, rgba(113, 113, 122, 0.15), rgba(82, 82, 91, 0.08));
    --fg: #fafafa; /* zinc-50 */
    --muted: #a1a1aa; /* zinc-400 */
    --card-bg: #18181b; /* zinc-900 */
    --surface-bg: #27272a; /* zinc-800 */
    --border: rgba(82, 82, 91, 0.3); /* zinc-600 with opacity */
    --shadow: 0 30px 90px rgba(9, 9, 11, 0.8); /* zinc-950 shadow */
    --surface-shadow: 0 22px 60px rgba(9, 9, 11, 0.5);
    --accent: #d4d4d8; /* zinc-300 */
    --accent-hover: #e4e4e7; /* zinc-200 */
    --button-bg: rgba(212, 212, 216, 0.1); /* zinc-300 with opacity */
    --button-border: rgba(161, 161, 170, 0.25); /* zinc-400 with opacity */
    --button-hover-bg: rgba(212, 212, 216, 0.2); /* zinc-300 with opacity */
    --avatar-border: rgba(161, 161, 170, 0.4); /* zinc-400 with opacity */
    --avatar-shadow: 0 10px 30px rgba(9, 9, 11, 0.25); /* zinc-950 shadow */
  }
}
body {
  margin: 0;
  padding: 0;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen-Sans, Ubuntu, Cantarell, "Helvetica Neue", sans-serif;
  background: var(--bg);
  background-image: var(--bg-gradient);
  color: var(--fg);
}
.page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 32px;
}
.container {
  width: min(520px, 100%);
}
.card {
  background: var(--card-bg);
  border-radius: 12px;
  border: 1px solid var(--border);
}
.shell {
  display: flex;
  flex-direction: column;
  gap: 28px;
  padding: 48px;
}
.surface {
  background: var(--surface-bg);
  border-radius: 8px;
  padding: 24px;
  border: 1px solid var(--border);
  box-shadow: var(--surface-shadow);
}
h1 {
  margin: 0;
  font-size: 2.3rem;
  line-height: 1.1;
}
p {
  margin: 0;
  color: var(--muted);
  line-height: 1.7;
}
.muted {
  color: var(--muted);
}
.surface > * + * {
  margin-top: 18px;
}
.data-row {
  display: flex;
  gap: 12px;
  align-items: baseline;
  font-weight: 500;
}
.data-label {
  min-width: 72px;
  font-size: 0.75rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: var(--muted);
}
.data-row strong {
  color: var(--fg);
  font-weight: 600;
}
.actions {
  display: flex;
  flex-wrap: wrap;
  gap: 16px;
}
.button {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 12px 20px;
  border-radius: 8px;
  border: 1px solid var(--button-border);
  background: var(--button-bg);
  color: var(--accent);
  font-weight: 600;
  text-decoration: none;
  transition: background 0.2s ease, transform 0.2s ease, border 0.2s ease;
}
.button:hover {
  background: var(--button-hover-bg);
  transform: translateY(-1px);
}
.button:active {
  transform: translateY(1px);
}
.avatar {
  display: flex;
  justify-content: center;
}
.avatar img {
  border-radius: 50%;
  width: 88px;
  height: 88px;
  border: 3px solid var(--avatar-border);
  box-shadow: var(--avatar-shadow);
}
@media (max-width: 640px) {
  .shell {
    padding: 32px;
  }
  .actions {
    flex-direction: column;
    align-items: stretch;
  }
}
`
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
