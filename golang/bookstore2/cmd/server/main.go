package main

import (
	"os"

	"bookstore2/core"
)

func main() {
	dbPath := "books.db"
	if len(os.Args) > 1 {
		dbPath = os.Args[1]
	}

	if err := core.RunServer(dbPath); err != nil {
		os.Exit(1)
	}
}
