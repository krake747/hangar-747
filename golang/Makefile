.PHONY: all build clean test gotest deps fmt fmt-md help

# Default target
all: deps fmt test build

# Build the CLI application
build:
	@echo "Building CLI..."
	go build -o bin/cli ./cmd/list

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	go clean

# Run tests
test:
	@echo "Running tests..."
	gotestdox -v ./...

# Download dependencies and install tools
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy
	@echo "Installing CLI tools..."
	@if ! command -v gofumpt >/dev/null 2>&1; then \
		go install mvdan.cc/gofumpt@latest; \
	fi
	@if ! command -v gotestdox >/dev/null 2>&1; then \
		go install github.com/bitfield/gotestdox/cmd/gotestdox@latest; \
	fi

# Format code
fmt:
	@echo "Formatting code..."
	gofumpt -l -w .

# Format markdown files
fmt-md:
	@echo "Formatting markdown files..."
	npx prettier --write '**/*.md'

# Show help
help:
	@echo "Available commands:"
	@echo "  build    - Build the CLI application"
	@echo "  clean    - Clean build artifacts"
	@echo "  test     - Run tests"
	@echo "  gotest   - Run tests with gotestdox"
	@echo "  deps     - Download dependencies and install tools"
	@echo "  fmt      - Format code"
	@echo "  fmt-md   - Format markdown files"
	@echo "  help     - Show this help message"
