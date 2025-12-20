# Bookstore

This was my starting point in building a Go web application, using simple file-based storage to
learn the fundamentals of HTTP servers and data management.

Built with basic net/http and JSON files, it taught me the basics of routing, concurrency with
`sync.RWMutex`, and simple API design.

## Features

You can manage a catalog of books with titles, authors, and stock levels. The RESTful API allows
listing, finding, and updating inventory. CLI tools provide easy interaction with the server.

## Getting Started

Run `make all` to install dependencies, test, and build everything. Start the server with
`./bin/server` (defaults to `localhost:8080` with `catalog` file). Use CLI tools like `./bin/list`
to interact.

## API & CLI

**CLIs**: server, list, find <ID>, add <ID> <Title> <Author>, setcopies <ID> <Count>. Connect to
localhost:3000 by default.
