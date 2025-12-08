package books

import (
	"fmt"
	"maps"
	"slices"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}

func GetCatalog() map[string]Book {
	catalog := map[string]Book{
		"abc": {ID: "abc", Title: "In the Company of Cheerful Ladies", Author: "Alexander McCall Smith", Copies: 1},
		"def": {ID: "def", Title: "White Heat", Author: "Dominic Sandbrook", Copies: 2},
	}
	return catalog
}

func BookToString(b Book) string {
	return fmt.Sprintf("%v by %v (copies: %v)", b.Title, b.Author, b.Copies)
}

func GetAllBooks(catalog map[string]Book) []Book {
	books := maps.Values(catalog)
	return slices.Collect(books)
}

func GetBook(catalog map[string]Book, ID string) (Book, bool) {
	book, ok := catalog[ID]
	return book, ok
}

func AddBook(catalog map[string]Book, b Book) {
	catalog[b.ID] = b
}
