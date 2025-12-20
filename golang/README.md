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

Focusing on translating C# skills to Go, I migrate the file-based bookstore to an Echo and SQLite
API. Echo's routing and middleware feel similar to .NET's minimal API and middleware integration,
but Go's explicit error handling contrasts with C#'s exceptions. Problem details for API errors
mirror .NET's structured responses.

File-based persistence lacks proper concurrency, transactions, and query performance – it's fine for
simple cases but falls short for real applications. Integrating SQLite is a step up since it's an
actual database, providing ACID transactions and efficient queries, much like working with a
relational database like PostgreSQL or SQL Server.

Structured logging with slog parallels Serilog, providing structured JSON output for better observability. Testing in Go is straightforward like xUnit, with built-in test runners and table-driven tests making it easy to cover edge cases.

Setting up golangci-lint feels like adding Roslyn analyzers in .NET, enforcing code quality automatically. GitHub Actions for CI/CD expands my knowledge beyond Azure Pipelines, showing how to automate builds, tests, and deployments in a cloud-native way.

Defer patterns for cleanup are new, replacing C#'s using statements. There's no dependency injection
here, unlike .NET's common DI containers. Overall, this project teaches me Go's simplicity and
efficiency, building on my C# foundations for maintainable services.
