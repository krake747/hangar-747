.PHONY: all build clean test gotest deps fmt fmt-md help

all: deps fmt test build

build:
	@echo "Building CLI..."
	go build -o bin/cli ./cmd/list

clean:
	@echo "Cleaning..."
	rm -rf bin/
	go clean

test:
	@echo "Running tests..."
	gotestdox -v ./...

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

fmt:
	@echo "Formatting code..."
	gofumpt -l -w .

fmt-md:
	@echo "Formatting markdown files..."
	npx prettier --write '**/*.md'

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
