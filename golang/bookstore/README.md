# Bookstore

A simple book inventory management system built with Go, using file-based storage.

## Features

- Manage a catalog of books with title, author, and stock levels.
- RESTful API for listing, finding, and updating book inventory.
- CLI tools for interacting with the server.

## Getting Started

Run `make all` to install dependencies, test, and build. Start the server with `./bin/server [catalog_path]` (defaults to `catalog` and `localhost:8080`). Use CLIs like `./bin/list` to interact.

## API Endpoints

- `GET /list`: List all books.
- `GET /find/{id}`: Find a book by ID.
- `POST /add/{id}/{title}/{author}`: Add a book.
- `POST /setcopies/{id}/{count}`: Set stock.

## CLI Tools

- `server`: Start the HTTP server.
- `list`: List all books.
- `find <ID>`: Find a book by ID.
- `add <ID> <Title> <Author>`: Add a book.
- `setcopies <ID> <Count>`: Set stock.

All CLIs connect to `localhost:8080` by default.