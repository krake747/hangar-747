package bookstore2_test

import (
	"cmp"
	"database/sql"
	"slices"
	"testing"

	"bookstore2/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetAllBooks_ReturnsAllBooks(t *testing.T) {
	t.Parallel()
	bs := getTestBookStore(t)
	bookList, err := bs.GetAllBooks()
	if err != nil {
		t.Fatal(err)
	}
	assertTestBooks(t, bookList)
}

func TestGetBook_FindsBookInBookStoreByID(t *testing.T) {
	t.Parallel()
	bs := getTestBookStore(t)
	got, err := bs.GetBook("abc")
	if err != nil {
		t.Fatal(err)
	}
	if got != ABC {
		t.Fatalf("want %#v, got %#v", ABC, got)
	}
}

func TestGetBook_ReturnsErrorWhenBookNotFound(t *testing.T) {
	t.Parallel()
	bs := getTestBookStore(t)
	_, err := bs.GetBook("nonexistent ID")
	if err == nil {
		t.Fatal("want error for nonexistent ID, got nil")
	}
}

func TestAddBook_AddsGivenBookToBookStore(t *testing.T) {
	t.Parallel()
	bs := getTestBookStore(t)
	_, err := bs.GetBook("123")
	if err == nil {
		t.Fatal("book already present")
	}
	err = bs.AddBook(models.Book{
		ID:     "123",
		Title:  "The Prize of all the Oceans",
		Author: "Glyn Williams",
		Copies: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = bs.GetBook("123")
	if err != nil {
		t.Fatal("added book not found")
	}
}

func TestSetCopies_SetsNumberOfCopiesToGivenValue(t *testing.T) {
	t.Parallel()
	book := models.Book{
		Copies: 5,
	}
	err := book.SetCopies(12)
	if err != nil {
		t.Fatal(err)
	}
	if book.Copies != 12 {
		t.Errorf("want 12 copies, got %d", book.Copies)
	}
}

func TestSetCopies_ReturnsErrorIfCopiesNegative(t *testing.T) {
	t.Parallel()
	book := models.Book{}
	err := book.SetCopies(-1)
	if err == nil {
		t.Error("want error for negative copies, got nil")
	}
}

func TestSetCopies_OnBookStoreModifiesSpecifiedBook(t *testing.T) {
	t.Parallel()
	bs := getTestBookStore(t)
	copies, err := bs.GetCopies("abc")
	if err != nil {
		t.Fatal(err)
	}
	if copies != 1 {
		t.Fatalf("want 1 copy before change, got %d", copies)
	}
	err = bs.SetCopies("abc", 2)
	if err != nil {
		t.Fatal(err)
	}
	copies, err = bs.GetCopies("abc")
	if err != nil {
		t.Fatal(err)
	}
	if copies != 2 {
		t.Fatalf("want 2 copies after change, got %d", copies)
	}
}

func getTestBookStore(t *testing.T) *models.BookStore {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	// Create table and insert test data
	_, err = db.Exec(`CREATE TABLE books (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		copies INTEGER NOT NULL DEFAULT 0
	)`)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(`INSERT INTO books (id, title, author, copies) VALUES
		('abc', 'In the Company of Cheerful Ladies', 'Alexander McCall Smith', 1),
		('xyz', 'White Heat', 'Dominic Sandbrook', 2)`)
	if err != nil {
		t.Fatal(err)
	}
	return models.NewBookStore(db)
}

func assertTestBooks(t *testing.T, got []models.Book) {
	t.Helper()
	want := []models.Book{ABC, XYZ}
	slices.SortFunc(got, func(a, b models.Book) int {
		return cmp.Compare(a.Author, b.Author)
	})
	if !slices.Equal(want, got) {
		t.Fatalf("want %#v, got %#v", want, got)
	}
}

var (
	ABC = models.Book{
		Title:  "In the Company of Cheerful Ladies",
		Author: "Alexander McCall Smith",
		Copies: 1,
		ID:     "abc",
	}

	XYZ = models.Book{
		Title:  "White Heat",
		Author: "Dominic Sandbrook",
		Copies: 2,
		ID:     "xyz",
	}
)
