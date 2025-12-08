package books_test

import (
	"books"
	"slices"
	"testing"
)

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
	got := books.GetAllBooks()
	if !slices.Equal(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestGetBook_FindsBookInCatalogByID(t *testing.T) {
	t.Parallel()
	want := books.Book{ID: "abc", Title: "In the Company of Cheerful Ladies", Author: "Alexander McCall Smith", Copies: 1}
	got, ok := books.GetBook("abc")
	if !ok {
		t.Fatalf("book not found")
	}
	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestGetBook_ReturnsFalseWhenBookNotFound(t *testing.T) {
	t.Parallel()
	_, ok := books.GetBook("nonexistent ID")
	if ok {
		t.Fatalf("want false for nonexistent book ID, got true")
	}
}
