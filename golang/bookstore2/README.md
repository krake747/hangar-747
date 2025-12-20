# Bookstore2

A simple book inventory management system built with Go, Echo, and SQLite.

This is a re-creation of the original bookstore application, modernized with a web framework and database persistence.

## Features

- Manage a catalog of books with title, author, and stock levels.
- RESTful API for listing, finding, and updating book inventory.
- CLI tools for interacting with the server.
- SQLite database for persistent storage.

## Getting Started

1. Install dependencies: `make deps`
2. Run tests: `make test`
3. Build CLIs: `make build`
4. Start the server: `./bin/server [db_path]`
   - Defaults to `books.db` and `localhost:3000`.
5. Use CLIs, e.g., `./bin/list` to view books.

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
- Build all: `make all`