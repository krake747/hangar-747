package main

import (
	"fmt"
	"os"

	"books"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: server <CATALOG FILE>")
		return
	}
	catalog, err := books.OpenCatalog(os.Args[1])
	if err != nil {
		fmt.Printf("opening catalog: %v\n", err)
		return
	}
	err = books.ListenAndServe("0.0.0.0:3000", catalog)
	if err != nil {
		fmt.Println(err)
	}
}
