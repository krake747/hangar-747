package books_test

import (
	"books"
	"cmp"
	"slices"
	"testing"
)

func getTestCatalog() map[string]books.Book {
	return map[string]books.Book{
		"abc": {ID: "abc", Title: "In the Company of Cheerful Ladies", Author: "Alexander McCall Smith", Copies: 1},
		"def": {ID: "def", Title: "White Heat", Author: "Dominic Sandbrook", Copies: 2},
	}
}

func TestBookToString_FormatsBookInfoAsString(t *testing.T) {
	t.Parallel()
	input := books.Book{
		Title:  "Sea Room",
		Author: "Adam Nicolson",
		Copies: 2,
	}
	want := "Sea Room by Adam Nicolson (copies: 2)"
	got := books.BookToString(input)
	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	t.Parallel()
	want := []books.Book{
		{ID: "abc", Title: "In the Company of Cheerful Ladies", Author: "Alexander McCall Smith", Copies: 1},
		{ID: "def", Title: "White Heat", Author: "Dominic Sandbrook", Copies: 2},
	}
	got := books.GetAllBooks(getTestCatalog())
	slices.SortFunc(got, func(a, b books.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})
	if !slices.Equal(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestGetBook_FindsBookInCatalogByID(t *testing.T) {
	t.Parallel()
	want := books.Book{ID: "abc", Title: "In the Company of Cheerful Ladies", Author: "Alexander McCall Smith", Copies: 1}
	got, ok := books.GetBook(getTestCatalog(), "abc")
	if !ok {
		t.Fatalf("book not found")
	}
	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()
	_, ok := books.GetBook(getTestCatalog(), "nonexistent ID")
	if ok {
		t.Fatalf("want false for nonexistent book ID, got true")
	}
}

func TestAddBook_AddsGiveBookToCatalog(t *testing.T) {
	t.Parallel()
	catalog := getTestCatalog()
	_, ok := books.GetBook(catalog, "123")
	if ok {
		t.Fatalf("book already present")
	}
	books.AddBook(catalog, books.Book{
		ID:     "123",
		Title:  "The Prize of all the Oceans",
		Author: "Glyn Williams",
		Copies: 2,
	})
	_, ok = books.GetBook(catalog, "123")
	if !ok {
		t.Fatalf("added book not found")
	}
}
