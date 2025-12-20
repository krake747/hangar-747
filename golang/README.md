# Go Learning Journey

This folder tracks my progress as I learn Go for cloud-native development and DevOps tooling.

## Projects

- [Bookstore](./bookstore/README.md) - Original file-based implementation
- [Bookstore2](./bookstore2/README.md) - Modernized with Echo and SQLite

## Learning Resources

- **[Tour of Go](https://go.dev/tour/)**
- **[Effective Go](https://go.dev/doc/effective_go)**
- **[The Deeper Love of Go](https://bitfieldconsulting.com/books/deeper)** (by Josh Arundel) - A
  good starter book; it's an excellent way to get into Go. The book store project Josh Arundel
  builds may be just simple CRUD, but that is the most common form of projects.

## 1. Learn Go Basics (Completed)

Focus areas:

- syntax, types, structs, functions, methods
- slices, maps, pointers
- error handling
- modules
- basic concurrency (goroutines, channels)

Goal: write small, idiomatic Go programs.

### Reflections

I learned the fundamentals of Go. The language does not provide safeguards, it is a very _explicit_
language which is the complete opposite of what I enjoy the most, _expressive_ functional languages.
However, going through the book I appreciate how simple it is to get started with Go.

Things I still need to wrap my head around are goroutines, since I am more familiar with the
async/await paradigm for concurrency.

There are still more advanced topics like channels, select statements, context, interfaces, and
generics that I need to experience, play around with, and learn.

I mostly just followed the examples from the book and re-implemented the bookstore API in this repo.
It is a good starting point for expansion.

Since Job Arundel is a big fan of testing, I added a GitHub Actions workflow that runs the tests via
`make` in CI. This is expanding my knowledge on DevOps with simple CI workflows.

## 2. Build a Simple Web API with Echo (Current)

Focus areas:

- set up an Echo project
- basic routing and middleware
- JSON responses and error handling
- database integration with SQLite
- structured logging and testing
- linting and CI/CD

Goal: create a modern Go web service with best practices.

### Reflections

I migrated the original file-based bookstore to a database-backed API using Echo, learning the power
of web frameworks for structured development. Echo simplified routing and middleware, allowing me to
focus on business logic rather than low-level HTTP handling. Integrating SQLite taught me about
transactions, prepared statements, and data persistence beyond files.

Adopting structured logging with slog improved observability, and adding comprehensive tests ensured
reliability. Setting up golangci-lint for code quality and GitHub Actions for CI expanded my DevOps
skills, enforcing standards automatically.

The migration highlighted the importance of modular architecture, separating handlers, models, and
clients. I also learned about defer patterns for resource management and problem details for API
errors. This project solidified my understanding of building production-ready Go services.

I expanded the original CRUD operations, added proper error responses, and integrated tooling for
maintainability. The journey from simple file storage to a robust API demonstrated Go's strengths in
web development.
