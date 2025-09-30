# PlainKit Starter

A minimal, production-ready starter template for building web applications with Go and PlainKit. This starter follows clean architecture principles and is designed to be familiar to developers coming from PHP, Python, or Node.js while remaining idiomatic Go.

## Features

- 🏗️ **Clean Architecture** - Separated concerns with domain, service, store, and handler layers
- 📦 **In-Memory Store** - Easy to swap for PostgreSQL, MySQL, or any database
- 🎨 **PlainKit HTML** - Type-safe HTML generation without templates
- 🌈 **Tailwind CSS v4** - Utility-first CSS with standalone CLI (no Node.js required)
- 🧪 **Integration Testing** - Testify-based tests focusing on real-world scenarios
- 🔄 **Hot Reload Ready** - Works with air, modd, or similar tools
- 📝 **Inline Comments** - Comprehensive documentation throughout

## Prerequisites

- **Go 1.25+** - [Download](https://go.dev/dl/)
- **Tailwind CSS CLI** - [Installation instructions below](#tailwind-css-installation)

## Quick Start

### 1. Install Dependencies

```bash
go mod download
```

### 2. Install Tailwind CSS CLI

```bash
# macOS
brew install tailwindcss

# Or download binary directly
# https://github.com/tailwindlabs/tailwindcss/releases
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
chmod +x tailwindcss-macos-arm64
mv tailwindcss-macos-arm64 /usr/local/bin/tailwindcss
```

### 3. Run the Application

```bash
# Using Makefile (builds CSS automatically)
make dev

# Or manually
tailwindcss -i ./internal/css/index.css -o ./internal/css/output.css --minify
go run cmd/server/main.go
```

### 4. Visit the Application

Open http://localhost:8080 in your browser.

## Project Structure

```
starter/
├── cmd/
│   └── server/
│       └── main.go              # Entry point (~30 lines)
│
├── internal/                    # Private application code
│   ├── app/
│   │   ├── app.go               # Dependency injection & routing
│   │   └── app_test.go          # Integration tests
│   │
│   ├── handlers/                # HTTP request handlers
│   │   ├── handlers.go          # Base handlers (home, health)
│   │   └── todos.go             # Todo-specific handlers
│   │
│   ├── domain/                  # Business entities
│   │   └── todo.go              # Todo entity
│   │
│   ├── service/                 # Business logic
│   │   ├── todo.go              # Todo service
│   │   └── todo_test.go         # Service tests
│   │
│   ├── store/                   # Data persistence
│   │   ├── memory.go            # In-memory store implementation
│   │   └── memory_test.go       # Store tests
│   │
│   ├── middleware/              # HTTP middleware
│   │   ├── logger.go            # Request logging
│   │   └── recover.go           # Panic recovery
│   │
│   ├── views/                   # HTML rendering
│   │   ├── layout.go            # Base layout
│   │   ├── home.go              # Home page
│   │   └── todos.go             # Todo list page
│   │
│   └── css/                     # Tailwind CSS
│       ├── embed.go             # Embeds output.css
│       ├── index.css            # Tailwind input with design tokens
│       └── output.css           # Generated (gitignored)
│
├── .env.example                 # Environment variables template
├── .gitignore
├── go.mod
├── Makefile                     # Build commands
└── README.md
```

## Available Commands

```bash
make help              # Show all commands
make dev               # Build CSS and run development server
make build             # Build CSS and compile binary
make test              # Run all tests
make test-integration  # Run integration tests only
make css               # Build CSS once (minified)
make css-watch         # Watch and rebuild CSS on changes
make clean             # Clean build artifacts
```

## Architecture Overview

This starter follows the structure defined in [GO_PROJECT_STRUCTURE.md](../../GO_PROJECT_STRUCTURE.md).

### Layers

1. **`cmd/server/main.go`** - Minimal entry point that bootstraps the app
2. **`internal/app/`** - Dependency injection and route configuration
3. **`internal/handlers/`** - HTTP handlers (parse requests, call services, render views)
4. **`internal/service/`** - Business logic and validation
5. **`internal/store/`** - Data persistence (interface + implementations)
6. **`internal/domain/`** - Core business entities (plain structs)
7. **`internal/middleware/`** - Cross-cutting concerns (logging, recovery)
8. **`internal/views/`** - HTML generation with PlainKit

### Data Flow

```
HTTP Request
    ↓
Middleware (logger, recover)
    ↓
Handler (parse request)
    ↓
Service (validate, business logic)
    ↓
Store (persistence)
    ↓
Service (return result)
    ↓
Handler (render view)
    ↓
HTTP Response
```

## Testing Strategy

This starter emphasizes **integration tests** over heavy unit testing:

- **Integration tests** in `internal/app/app_test.go` test complete workflows
- **Unit tests** for complex business logic in `internal/service/`
- **Unit tests** for store implementations in `internal/store/`
- Uses [Testify](https://github.com/stretchr/testify) for assertions

### Running Tests

```bash
# All tests
make test

# Integration tests only
make test-integration

# Specific test
go test -v ./internal/service -run TestTodoService_Create

# With coverage
go test -cover ./...
```

## Extending the Starter

### Adding a New Feature

1. **Define the entity** in `internal/domain/`
2. **Create store interface and implementation** in `internal/store/`
3. **Add business logic** in `internal/service/`
4. **Create handlers** in `internal/handlers/`
5. **Add views** in `internal/views/`
6. **Register routes** in `internal/app/app.go`
7. **Write tests** colocated with code

### Swapping the Data Store

Replace `MemoryStore` with a real database:

```go
// internal/store/postgres.go
type PostgresStore struct {
    db *sql.DB
}

func NewPostgresStore(db *sql.DB) *PostgresStore {
    return &PostgresStore{db: db}
}

func (s *PostgresStore) Create(ctx context.Context, todo *domain.Todo) error {
    _, err := s.db.ExecContext(ctx,
        "INSERT INTO todos (id, title, completed) VALUES ($1, $2, $3)",
        todo.ID, todo.Title, todo.Completed,
    )
    return err
}
// ... implement other methods
```

Then update `internal/app/app.go`:

```go
func New() *App {
    db := connectDB() // your DB connection logic
    todoStore := store.NewPostgresStore(db)
    // ... rest of initialization
}
```

### Adding Middleware

Create a new file in `internal/middleware/`:

```go
// internal/middleware/cors.go
func CORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        next.ServeHTTP(w, r)
    })
}
```

Apply it in `internal/app/app.go`:

```go
return middleware.Logger(
    middleware.CORS(
        middleware.Recover(mux),
    ),
)
```

### Adding UI Components

Copy PlainUI components from the main library as needed:

```bash
# Manually copy components
cp -r ../../ui/button internal/ui/
cp -r ../../ui/card internal/ui/

# Future: CLI tool will automate this
plainui add button card
```

## Development Tips

### Watch Mode

Run with auto-reload using two terminals:

```bash
# Terminal 1: Watch CSS
make css-watch

# Terminal 2: Run server with hot reload
# Install: go install github.com/cosmtrek/air@latest
air
```

### Environment Variables

Copy `.env.example` to `.env` and customize:

```bash
cp .env.example .env
```

The app uses Bun's automatic `.env` loading (no library needed).

### Debugging

Add logging anywhere:

```go
log.Printf("debug: value=%v", someValue)
```

## Deployment

### Build for Production

```bash
make build
```

This creates a single binary at `bin/server` with embedded CSS.

### Docker

```dockerfile
FROM golang:1.25-alpine AS builder
RUN apk add --no-cache tailwindcss
WORKDIR /app
COPY . .
RUN make build

FROM alpine:latest
COPY --from=builder /app/bin/server /server
EXPOSE 8080
CMD ["/server"]
```

### Environment Variables

Set these in production:

- `PORT` - Server port (default: 8080)
- Add your own as needed (database URLs, API keys, etc.)

## For Developers from Other Ecosystems

### Coming from Node.js/Express?

| Express             | Go Equivalent                |
|---------------------|------------------------------|
| `app.js`            | `cmd/server/main.go`         |
| `routes/`           | `internal/app/app.go`        |
| `controllers/`      | `internal/handlers/`         |
| `services/`         | `internal/service/`          |
| `models/`           | `internal/domain/`           |
| `repositories/`     | `internal/store/`            |
| `middlewares/`      | `internal/middleware/`       |
| `views/`            | `internal/views/`            |

### Coming from Python/Django?

| Django              | Go Equivalent                |
|---------------------|------------------------------|
| `manage.py`         | `cmd/server/main.go`         |
| `urls.py`           | `internal/app/app.go`        |
| `views.py`          | `internal/handlers/`         |
| `models.py`         | `internal/domain/`           |
| `services.py`       | `internal/service/`          |
| `middleware.py`     | `internal/middleware/`       |
| `templates/`        | `internal/views/`            |

### Coming from PHP/Laravel?

| Laravel                  | Go Equivalent                |
|--------------------------|------------------------------|
| `public/index.php`       | `cmd/server/main.go`         |
| `routes/web.php`         | `internal/app/app.go`        |
| `app/Http/Controllers/`  | `internal/handlers/`         |
| `app/Services/`          | `internal/service/`          |
| `app/Models/`            | `internal/domain/`           |
| `app/Repositories/`      | `internal/store/`            |
| `app/Http/Middleware/`   | `internal/middleware/`       |
| `resources/views/`       | `internal/views/`            |

## Additional Resources

- [PlainKit HTML](https://github.com/plainkit/html) - Type-safe HTML generation
- [PlainKit HTMX](https://github.com/plainkit/htmx) - HTMX integration
- [Tailwind CSS](https://tailwindcss.com/) - Utility-first CSS
- [Testify](https://github.com/stretchr/testify) - Testing toolkit
- [Go Project Structure Guide](../../GO_PROJECT_STRUCTURE.md) - Detailed architecture guide

## License

MIT