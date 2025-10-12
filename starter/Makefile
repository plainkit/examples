.PHONY: help css css-watch dev build test clean

help: ## Show available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

css: ## Compile Tailwind CSS once
	tailwindcss -i ./internal/css/index.css -o ./internal/css/output.css --minify

css-watch: ## Rebuild CSS on file changes
	tailwindcss -i ./internal/css/index.css -o ./internal/css/output.css --watch

dev: css ## Build CSS and run the development server
	go run ./cmd/server

build: css ## Build the starter binary
	go build -o ./bin/starter ./cmd/server

test: ## Run unit and integration tests
	go test ./...

clean: ## Remove build artifacts
	rm -rf ./bin ./internal/css/output.css
