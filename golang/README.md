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

## 2. Build a Simple Web API with Echo (Completed)

Focus areas:

- set up an [Echo](https://echo.labstack.com/) project
- basic routing and middleware
- JSON responses and error handling
- database integration with SQLite
- structured logging and testing
- linting and CI/CD

Goal: create a modern Go web service with best practices.

### Reflections

This step focuses on applying my C# experience in a Go project and understanding which skills
transfer directly and which require a different mindset. Rebuilding the file-based bookstore as an
Echo API with SQLite helps me focus less on the language and more on service design.

Echo feels familiar. Routing and middleware are close to .NET’s minimal APIs, which makes it easy to
get productive. Go’s error handling forces me to be more explicit. Without exceptions, I have to
think through error paths more carefully, which makes the code easier to reason about and the API
behavior more predictable. Using problem details for errors aligns well with the API patterns I
already use in .NET.

The limitations of file-based persistence become clear quickly. Concurrency, transactions, and
querying are hard to get right without a real database. Switching to SQLite is an important step.
Even though it is lightweight, it provides ACID guarantees and behaves like a proper relational
database, similar to PostgreSQL or SQL Server.

I also focus more on operational concerns. Structured logging with `slog` helps me think about logs
as structured data and prepares the codebase for adding OpenTelemetry later. Testing in Go is
straightforward, and table-driven tests encourage simple and readable test cases.

Adding `golangci-lint` and GitHub Actions brings the project closer to production standards.
Automating tests and linting in CI reinforces good habits and expands my CI/CD experience beyond
Azure Pipelines.

Go does not rely on dependency injection containers like .NET. Wiring dependencies manually feels
different. It makes data flow and ownership very clear and easy to understand.

Overall, this step helps me build confidence in designing maintainable services in Go while applying
the same engineering principles I use in my .NET work.

## 3. Build a Production-Ready API with Go (Planned)

This step focuses on evolving the bookstore API toward a more production-ready, cloud-native
service. Key areas of focus include:

- Database -> PostgreSQL or MongoDB (TBD)
- Containerization -> Dockerfile, Docker Compose
- Observability -> OpenTelemetry
- Configuration Management -> Environment variables, config files
- Secrets -> Vault
