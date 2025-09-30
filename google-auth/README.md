# PlainKit + Goth Google OAuth Example

A minimal PlainKit HTML web app that demonstrates Google OAuth authentication using [goth](https://github.com/markbates/goth).

## Features

- 🔐 Google OAuth 2.0 authentication
- 🎨 PlainUI component library integration
- 📦 Clean architecture with separated concerns
- 🔄 Session-based authentication
- 🎯 Protected routes with middleware

## Prerequisites

- Go 1.25+
- [Tailwind CSS CLI](https://tailwindcss.com/docs/installation) (standalone binary, no Node.js required)
- A Google Cloud project with OAuth credentials (web application)

## Google OAuth Setup

1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project or select an existing one
3. Navigate to **APIs & Services** > **Credentials**
4. Create **OAuth 2.0 Client ID** credentials:
   - Application type: **Web application**
   - Authorized redirect URI: `http://localhost:3000/auth/google/callback`
5. Copy the **Client ID** and **Client Secret**

## Quick Start

1. Clone and navigate to the directory:
   ```bash
   cd examples/google-auth
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Install Tailwind CSS CLI:
   ```bash
   # macOS
   brew install tailwindcss

   # Or download from GitHub releases
   # https://github.com/tailwindlabs/tailwindcss/releases
   ```

4. Copy environment file and add your credentials:
   ```bash
   cp .env.example .env
   # Edit .env with your Google OAuth credentials
   ```

5. Build CSS and run the server:
   ```bash
   # Using Makefile
   make dev

   # Or manually
   tailwindcss -i ./internal/css/index.css -o ./internal/css/output.css --minify
   go run cmd/server/main.go
   ```

6. Visit http://localhost:3000 and click **Continue with Google**

## Project Structure

```
cmd/server/           - Application entry point
internal/
  app/                - App initialization & routing
  handlers/           - HTTP request handlers
  domain/             - Business entities (User)
  middleware/         - HTTP middleware (auth)
  views/              - HTML templates
  ui/                 - PlainUI components (copied from /ui)
  icons/              - Icon components
  css/                - Tailwind CSS (index.css + embed.go)
.env.example          - Environment variables template
Makefile              - Build/dev commands
```

## Available Commands

```bash
make help         # Show all commands
make dev          # Build CSS and run development server
make build        # Build CSS and compile binary
make css          # Build CSS once (minified)
make css-watch    # Watch and rebuild CSS on changes
make test         # Run tests
make clean        # Clean build artifacts
```

## Environment Variables

Create a `.env` file based on `.env.example`:

- `GOOGLE_CLIENT_ID` - Your Google OAuth client ID **(required)**
- `GOOGLE_CLIENT_SECRET` - Your Google OAuth client secret **(required)**
- `BASE_URL` - Application base URL (default: `http://localhost:3000`)
- `PORT` - Server port (default: `3000`)
- `SESSION_SECRET` - Session encryption key (generate with `openssl rand -hex 32`)

## How It Works

1. **Home Page** (`/`) - Shows Google login button
2. **OAuth Flow** (`/auth/google`) - Redirects to Google for authentication
3. **Callback** (`/auth/google/callback`) - Handles OAuth callback, creates session
4. **Dashboard** (`/dashboard`) - Protected route showing user profile
5. **Logout** (`/logout`) - Clears session and redirects home

## Architecture

This example follows the structure recommended in [GO_PROJECT_STRUCTURE.md](../../GO_PROJECT_STRUCTURE.md):

- **`cmd/server/main.go`** - Minimal entry point (~30 lines)
- **`internal/app/`** - Dependency injection and route configuration
- **`internal/handlers/`** - HTTP handlers for pages and auth
- **`internal/domain/`** - User entity definition
- **`internal/middleware/`** - Authentication middleware
- **`internal/views/`** - HTML page rendering with PlainKit
- **`internal/ui/`** - PlainUI components (button, card, avatar, separator)
- **`internal/css/`** - Tailwind CSS with embedded output

## Development

### Watch Mode

Run with auto-reload using two terminals:

```bash
# Terminal 1: Watch CSS
make css-watch

# Terminal 2: Run server (use air or modd for auto-reload)
go run cmd/server/main.go
```

### Adding UI Components

PlainUI components are manually copied from `/ui` directory:

```bash
# Copy component from main UI library
cp -r ../../ui/badge internal/ui/

# Future: CLI tool will automate this
plainui add badge
```

## Testing

```bash
# Run all tests
make test

# Run specific test
go test -v ./internal/handlers -run TestHandlers
```

## Deployment

Build the application:

```bash
make build
```

This creates a single binary at `bin/server` with all assets embedded.

Deploy environment variables:
- Set `GOOGLE_CLIENT_ID` and `GOOGLE_CLIENT_SECRET`
- Set `BASE_URL` to your production domain
- Generate secure `SESSION_SECRET`
- Update Google OAuth redirect URI to match production URL

## License

MIT