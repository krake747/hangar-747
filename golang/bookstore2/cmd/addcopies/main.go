package main

import (
	"fmt"
	"os"
	"strconv"

	"bookstore2/internal/client"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: addcopies <BOOK_ID> <HOW MANY>")
		return
	}
	ID := os.Args[1]
	copies, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	c := client.NewClient("localhost:3000")
	stock, err := c.AddCopies(ID, copies)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d copies in stock", stock)
}
