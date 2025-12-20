package bookstore2

import (
	"database/sql"
	"log/slog"
	"net/http/httptest"
	"os"
	"testing"

	"bookstore2/core"
	"bookstore2/core/books"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

var ABC = books.Book{ID: "abc", Title: "In the Company of Cheerful Ladies", Author: "Alexander McCall Smith", Copies: 1}

func getTestClient(t *testing.T) *core.Client {
	t.Helper()

	// Init DB
	db := initTestDB(t)
	t.Cleanup(func() { db.Close() })

	bookStore := books.NewBookStore(db)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	h := books.NewHandler(logger, bookStore)

	// Setup Echo
	e := echo.New()
	e.GET("/v1/list", h.ListBooks)
	e.GET("/v1/find/:id", h.FindBook)
	e.GET("/v1/getcopies/:id", h.GetCopies)
	e.POST("/v1/addcopies/:id/:count", h.AddCopies)
	e.POST("/v1/subcopies/:id/:count", h.SubCopies)

	// Start test server
	ts := httptest.NewServer(e)
	t.Cleanup(ts.Close)

	// Create client with test server URL
	client := core.NewClient(ts.URL[7:]) // Remove http://

	return client
}

func initTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	initSQL, err := os.ReadFile("testdata/init.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(string(initSQL))
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func TestGetAllBooks_OnClientReturnsAllBooks(t *testing.T) {
	t.Parallel()
	client := getTestClient(t)
	bookList, err := client.GetAllBooks()
	if err != nil {
		t.Fatal(err)
	}
	expected := []books.Book{
		{ID: "abc", Title: "In the Company of Cheerful Ladies", Author: "Alexander McCall Smith", Copies: 1},
		{ID: "xyz", Title: "White Heat", Author: "Dominic Sandbrook", Copies: 2},
	}
	if len(bookList) != len(expected) {
		t.Fatalf("want %d books, got %d", len(expected), len(bookList))
	}
	for _, b := range bookList {
		if b.ID == "abc" && b.Copies != 1 {
			t.Errorf("abc copies want 1, got %d", b.Copies)
		}
		if b.ID == "xyz" && b.Copies != 2 {
			t.Errorf("xyz copies want 2, got %d", b.Copies)
		}
	}
}

func TestGetBook_OnClientFindsBookByID(t *testing.T) {
	t.Parallel()
	client := getTestClient(t)
	got, err := client.GetBook("abc")
	if err != nil {
		t.Fatal(err)
	}
	if got != ABC {
		t.Fatalf("want %#v, got %#v", ABC, got)
	}
}

func TestGetBook_OnClientReturnsErrorWhenBookNotFound(t *testing.T) {
	t.Parallel()
	client := getTestClient(t)
	_, err := client.GetBook("bogus")
	if err == nil {
		t.Error("want error when book not found, got nil")
	}
}

func TestGetCopies_OnClientReturnsCopiesForBook(t *testing.T) {
	t.Parallel()
	client := getTestClient(t)
	copies, err := client.GetCopies("abc")
	if err != nil {
		t.Fatal(err)
	}
	if copies != 1 {
		t.Fatalf("want 1 copy, got %d", copies)
	}
}

func TestGetCopies_OnClientErrorsWhenBookNotFound(t *testing.T) {
	t.Parallel()
	client := getTestClient(t)
	_, err := client.GetCopies("bogus")
	if err == nil {
		t.Error("want error when book not found, got nil")
	}
}
