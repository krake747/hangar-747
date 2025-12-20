.PHONY: all build clean test gotest deps fmt fmt-md help

all: deps fmt test build

build:
	@echo "Building CLIs..."
	mkdir -p bin
	go build -o bin/server ./cmd/server
	go build -o bin/list ./cmd/list
	go build -o bin/find ./cmd/find
	go build -o bin/getcopies ./cmd/getcopies
	go build -o bin/addcopies ./cmd/addcopies
	go build -o bin/subcopies ./cmd/subcopies
	go build -o bin/client ./cmd/client

clean:
	@echo "Cleaning..."
	rm -rf bin/
	go clean

test:
	@echo "Running tests..."
	go test -v ./...

deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy
	@echo "Installing CLI tools..."
	@if ! command -v gofumpt >/dev/null 2>&1; then \
		go install mvdan.cc/gofumpt@latest; \
	fi

fmt:
	@echo "Formatting code..."
	gofumpt -l -w .

fmt-md:
	@echo "Formatting markdown files..."
	npx prettier --write '**/*.md'

help:
	@echo "Available commands:"
	@echo "  build    - Build all CLI applications"
	@echo "  clean    - Clean build artifacts"
	@echo "  test     - Run tests"
	@echo "  deps     - Download dependencies and install tools"
	@echo "  fmt      - Format code"
	@echo "  fmt-md   - Format markdown files"
	@echo "  help     - Show this help message"