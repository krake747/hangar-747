package main

import (
	"fmt"
	"os"

	"bookstore2/core"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: find <BOOK ID>")
		return
	}
	ID := os.Args[1]
	c := core.NewClient("localhost:3000")
	book, err := c.GetBook(ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(book)
}
