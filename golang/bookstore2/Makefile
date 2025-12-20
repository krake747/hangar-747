.PHONY: all build clean test gotest deps fmt lint fmt-md help

all: deps fmt lint test build

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
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi

lint:
	@echo "Running linter..."
	$(shell go env GOPATH)/bin/golangci-lint run ./...

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
	@echo "  lint     - Run linter"
	@echo "  deps     - Download dependencies and install tools"
	@echo "  fmt      - Format code"
	@echo "  fmt-md   - Format markdown files"
	@echo "  help     - Show this help message"