# Bookstore2

A demonstration of migrating a simple file-based Go application to a modern web service using Echo
and SQLite, learning best practices along the way.

Originally a basic file-based bookstore with manual HTTP handling, bookstore2 showcases improvements
in persistence, error handling, and developer tooling through a hands-on migration.

## Migration Story

The original bookstore used a JSON file for storage with in-memory operations and basic net/http for
serving. This migration introduces:

- **Database Persistence**: SQLite for reliable, concurrent data handling instead of file locking.
- **Web Framework**: Echo for structured routing, middleware, and request management.
- **Quality Tools**: Linting, formatting, and CI/CD to enforce code standards.
- **Architecture**: Modular design with clear separation of HTTP handlers, business logic, and API
  clients.

## Features

Through the migration, we added:

- RESTful API with JSON responses and consistent error handling.
- Middleware for problem details (RFC 7807) to improve API usability.
- Structured logging with slog for better observability.
- SQLite database with transactions for data integrity.

## Getting Started

Run `make all` to install dependencies, lint, test, and build everything. Start the server with
`./bin/server` (defaults to `localhost:3000` with `books.db`). Use CLI tools like `./bin/list` to
interact.

## API & CLI

**CLIs**:

- server
- list
- find <ID>
- getcopies <ID>
- addcopies <ID> <count>
- subcopies <ID> <count>
- client

Connect to localhost:3000 by default.
