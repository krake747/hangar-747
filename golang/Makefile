.PHONY: all build clean test gotest deps fmt help

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
	go test -v ./...

gotest:
	@echo "Running tests with gotestdox..."
	go run github.com/bitfield/gotestdox/cmd/gotestdox@latest -v ./...

# Download dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Format code
fmt:
	@echo "Formatting code..."
	gofumpt -l -w .

# Show help
help:
	@echo "Available commands:"
	@echo "  build    - Build the CLI application"
	@echo "  clean    - Clean build artifacts"
	@echo "  test     - Run tests"
	@echo "  gotest   - Run tests with gotestdox"
	@echo "  deps     - Download dependencies"
	@echo "  fmt      - Format code"
	@echo "  help     - Show this help message"
