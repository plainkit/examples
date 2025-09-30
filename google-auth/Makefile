.PHONY: help dev build test clean css css-watch

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

css: ## Build Tailwind CSS (minified)
	@echo "Building CSS..."
	tailwindcss -i ./internal/css/index.css -o ./internal/css/output.css --minify

css-watch: ## Watch and rebuild CSS on changes
	tailwindcss -i ./internal/css/index.css -o ./internal/css/output.css --watch

dev: css ## Build CSS and run development server
	@echo "Starting server..."
	go run cmd/server/main.go

build: css ## Build CSS and compile binary
	@echo "Building application..."
	go build -o bin/server cmd/server/main.go

test: ## Run all tests
	go test -v -race -cover ./...

clean: ## Clean build artifacts
	rm -rf bin/ internal/css/output.css