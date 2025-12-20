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

Focusing on translating C# skills to Go, I migrate the file-based bookstore to an Echo and SQLite API. This highlights the differences: Echo's routing feels similar to ASP.NET middleware, but Go's explicit error handling contrasts with C#'s exceptions.

Integrating SQLite reminds me of ADO.NET and Dapper (since I know those), but Go's sql package requires more manual work with prepared statements and transactions – a trade-off for performance and control compared to Entity Framework.

Structured logging with slog parallels Serilog, and testing in Go is straightforward like xUnit. Setting up golangci-lint and GitHub Actions expands my CI knowledge, automating quality like in Azure DevOps.

Defer patterns for cleanup are new, replacing C#'s using statements. Overall, this project teaches me Go's simplicity and efficiency, building on my C# foundations for maintainable services.
