# Bookstore2

A demonstration of migrating a simple file-based Go application to a modern web service using Echo and SQLite. I learn best practices along the way.

Originally a basic file-based bookstore with manual HTTP handling, bookstore2 showcases improvements in persistence, error handling, and developer tooling through a hands-on migration.

## Migration Story

The original bookstore uses a JSON file for storage with in-memory operations and basic `net/http` for serving. Migrating it teaches me a lot about building maintainable web services.

Along the way, I tackle challenges like safely handling concurrent writes, structuring routes, and designing clear error responses. I introduce:

* **Database Persistence**: We move to SQLite to handle concurrent data reliably, replacing fragile file locks.
* **Web Framework**: Echo structures our routes, middleware, and request flow, letting me focus on features instead of plumbing.
* **Quality Tools**: I integrate linting, formatting, and CI/CD to keep the codebase clean and consistent.
* **Architecture**: I modularize the project, separating HTTP handlers, business logic, and API clients. This makes it easier to reason about and extend.

Every step feels like learning to think in terms of *robust, production-ready services*, not just toy examples. I experiment with structured logging, transactions, and problem-detail middleware, seeing immediate improvements in observability and reliability.

## Features

Through the migration, we add:

* RESTful API with JSON responses and consistent error handling.
* Middleware for problem details (RFC 7807) to improve API usability.
* Structured logging with `slog` for better observability.
* SQLite database with transactions for data integrity.

## Getting Started

I run `make all` to install dependencies, lint, test, and build everything. I start the server with `./bin/server` (defaulting to `localhost:3000` with `books.db`). I use CLI tools like `./bin/list` to interact.

## API & CLI

**CLIs**:

* `server`
* `list`
* `find <ID>`
* `getcopies <ID>`
* `addcopies <ID> <count>`
* `subcopies <ID> <count>`
* `client`

By default, we connect to `localhost:3000`.
