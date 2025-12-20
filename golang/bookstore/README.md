# Bookstore

This is my starting point in building a Go web application. I use simple file-based storage to learn
the fundamentals of HTTP servers and data management.

Built with basic `net/http` and JSON files, I learn the basics of routing, concurrency with
`sync.RWMutex`, and simple API design.

## Features

We manage a catalog of books with titles, authors, and stock levels. The RESTful API lets us list,
find, and update inventory. CLI tools let us interact with the server easily.

## Getting Started

I run `make all` to install dependencies, test, and build everything. I start the server with
`./bin/server` (defaulting to `localhost:3000` with the `catalog` file). I use CLI tools like
`./bin/list` to interact with it.

## API & CLI

**CLIs**:

- `server`
- `list`
- `find <ID>`
- `getcopies <ID>`
- `addcopies <ID> <count>`
- `subcopies <ID> <count>`
- `client`

By default, we connect to `localhost:3000`.
