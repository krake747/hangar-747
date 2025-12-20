package main

import (
	"fmt"

	"bookstore2/core"
)

func main() {
	c := core.NewClient("localhost:3000")
	bookList, err := c.GetAllBooks()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, book := range bookList {
		fmt.Println(book)
	}
}
