package main

import (
	"fmt"

	"bookstore2/internal/client"
)

func main() {
	c := client.NewClient("localhost:3000")
	bookList, err := c.GetAllBooks()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, book := range bookList {
		fmt.Println(book)
	}
}
