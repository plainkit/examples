# Starter Todo App

A monochrome, modern todo tracker showcasing how to build a Plainkit-powered Go web app with zero external services. It ships with in-memory storage, Pure Go HTML rendering, and embedded styles for instant previews.

## Features

- Sleek black-and-white UI with frosted cards and glowing CTAs
- In-memory persistence layer with domain/service separation
- Type-safe HTML markup using [plainkit/html](https://github.com/plainkit/html)
- Middleware stack for logging and panic recovery
- Integration and unit tests covering the happy path

## Quick Start

```bash
# From repo root
cd examples/starter

# Install Go dependencies
go mod tidy

# (Recommended) Install Tailwind CLI once
brew install tailwindcss # or download from the Tailwind releases page

# Compile CSS and run the server
make dev
```

Visit `http://localhost:8080` to start adding tasks.

## Tailwind Workflow

The UI is styled via Tailwind tokens defined in `internal/css/index.css`. To regenerate the embedded stylesheet, run:

```bash
make css
# or for live editing
tailwindcss -i ./internal/css/index.css -o ./internal/css/output.css --watch
```

`internal/css/output.css` is embedded into the Go binary (`internal/css/embed.go`), so remember to rebuild after recompiling CSS.

## Project Layout

```
cmd/server/           Entry point (minimal bootstrap)
internal/app/         Wiring + HTTP mux
internal/handlers/    HTTP handlers
internal/service/     Business logic for todos
internal/store/       In-memory persistence
internal/views/       Plainkit HTML layout + pages
internal/css/         Tailwind source + embedded output
testdata/             Fixture playground (empty by default)
```

## Testing

```bash
make test
```

- `internal/service/todo_test.go` covers the service layer
- `internal/app/app_test.go` runs an end-to-end todo flow against `httptest`

## Environment

Copy `.env.example` to `.env` if you want to override defaults (currently just `PORT`).

## Design Notes

- Typography: Inter, high-contrast sizing, tight leading
- Contrast: Gradients and glassmorphism lighten the deep black background
- Interactions: Subtle hover+shadow states on tasks and primary controls

Have fun hacking on it—swap the in-memory store for Postgres, add HTMX sprinkles, or port the design tokens into your next PlainUI project.
