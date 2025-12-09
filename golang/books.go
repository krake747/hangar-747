package books

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}

type Catalog map[string]Book

func OpenCatalog(path string) (Catalog, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	catalog := Catalog{}
	err = json.NewDecoder(file).Decode(&catalog)
	if err != nil {
		return nil, err
	}
	return catalog, nil
}

func (catalog Catalog) Sync(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(catalog)
	if err != nil {
		return err
	}
	return nil
}

func (catalog Catalog) SetCopies(ID string, copies int) error {
	book, ok := catalog[ID]
	if !ok {
		return fmt.Errorf("ID %q not found", ID)
	}
	err := book.SetCopies(copies)
	if err != nil {
		return err
	}
	catalog[ID] = book
	return nil
}

func GetCatalog() Catalog {
	catalog := map[string]Book{
		"abc": {ID: "abc", Title: "In the Company of Cheerful Ladies", Author: "Alexander McCall Smith", Copies: 1},
		"def": {ID: "def", Title: "White Heat", Author: "Dominic Sandbrook", Copies: 2},
	}
	return catalog
}

func (b Book) String() string {
	return fmt.Sprintf("%v by %v (copies: %v)", b.Title, b.Author, b.Copies)
}

func (b *Book) SetCopies(copies int) error {
	if copies < 0 {
		return fmt.Errorf("negative number of copies: %d", copies)
	}
	b.Copies = copies
	return nil
}

func (catalog Catalog) GetAllBooks() []Book {
	return slices.Collect(maps.Values(catalog))
}

func (catalog Catalog) GetBook(ID string) (Book, bool) {
	book, ok := catalog[ID]
	return book, ok
}

func (catalog Catalog) AddBook(book Book) error {
	_, ok := catalog[book.ID]
	if ok {
		return fmt.Errorf("ID %q already exists", book.ID)
	}
	catalog[book.ID] = book
	return nil
}
