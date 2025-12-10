# Go Learning Journey

This folder tracks my progress as I learn Go for cloud-native development and DevOps tooling.

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

Next step:

- set up an Echo project
- basic routing
- JSON responses
- request handling
- logging and middleware
- containerizing the API

Goal: create a minimal Go web service to build on later.
