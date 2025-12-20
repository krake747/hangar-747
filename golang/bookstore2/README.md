# Bookstore2

A demonstration of migrating a simple file-based Go application to a modern web service using Echo and SQLite, learning best practices along the way.

Originally a basic file-based bookstore with manual HTTP handling, bookstore2 showcases improvements in persistence, error handling, and developer tooling through a hands-on migration.

## Migration Story

The original bookstore used a JSON file for storage with in-memory operations and basic net/http for serving. This migration introduces:
- **Database Persistence**: SQLite for reliable, concurrent data handling instead of file locking.
- **Web Framework**: Echo for structured routing, middleware, and request management.
- **Quality Tools**: Linting, formatting, and CI/CD to enforce code standards.
- **Architecture**: Modular design with clear separation of HTTP handlers, business logic, and API clients.

## Features

Through the migration, we added:
- RESTful API with JSON responses and consistent error handling.
- Middleware for problem details (RFC 7807) to improve API usability.
- Structured logging with slog for better observability.
- Comprehensive test suite with parallel execution.
- CLI tools for all inventory operations.
- SQLite database with transactions for data integrity.

## Learnings from Migrating to Echo

- **Routing**: Echo's router simplifies defining and matching API endpoints compared to manual path parsing.
- **Middleware**: Enables centralized handling of cross-cutting concerns like errors, logging, and request processing.
- **Context Management**: Using echo.Context provides a clean way to handle request lifecycles, validation, and responses.
- **Error Handling**: Structured problem details make API errors more informative and consistent for clients.
- **Testing**: Echo's testing utilities make it easier to write robust integration tests for HTTP endpoints.
- **Persistence**: Database transactions ensure data consistency and handle concurrency better than file-based approaches.

## Getting Started

1. Install dependencies: `make deps`
2. Run tests and linter: `make test && make lint`
3. Build CLIs: `make build`
4. Start the server: `./bin/server [db_path]`
    - Defaults to `books.db` and `localhost:3000`.
5. Use CLIs, e.g., `./bin/list` to view books.

## API Endpoints

- `GET /v1/list`: List all books.
- `GET /v1/find/{id}`: Find a book by ID.
- `GET /v1/getcopies/{id}`: Get stock for a book.
- `POST /v1/addcopies/{id}/{count}`: Add stock.
- `POST /v1/subcopies/{id}/{count}`: Subtract stock.

All endpoints return JSON; errors use problem details format.

## CLI Tools

- `server`: Start the HTTP server.
- `list`: List all books.
- `find <ID>`: Find a book by ID.
- `getcopies <ID>`: Get stock for a book.
- `addcopies <ID> <count>`: Add stock.
- `subcopies <ID> <count>`: Subtract stock.
- `client`: Test server connectivity.

All CLIs connect to `localhost:3000` by default.

## Development

- Format code: `make fmt`
- Run tests: `make test`
- Lint code: `make lint`
- Build all: `make all`

Learned to integrate linting and formatting into the workflow for consistent code quality.