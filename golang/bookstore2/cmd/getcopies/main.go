package main

import (
	"fmt"
	"os"

	"bookstore2/core"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: getcopies <BOOK_ID>")
		return
	}
	ID := os.Args[1]
	c := core.NewClient("localhost:3000")
	copies, err := c.GetCopies(ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d copies in stock", copies)
}
