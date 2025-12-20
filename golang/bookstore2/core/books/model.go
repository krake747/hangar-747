package books

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ID     string
	Title  string
	Author string
	Copies int
}

func (book Book) String() string {
	return fmt.Sprintf("%v by %v (copies: %v)",
		book.Title, book.Author, book.Copies)
}

func (book *Book) SetCopies(copies int) error {
	if copies < 0 {
		return fmt.Errorf("negative number of copies: %d", copies)
	}
	book.Copies = copies
	return nil
}

type BookStore struct {
	DB *sql.DB
}

func NewBookStore(db *sql.DB) *BookStore {
	return &BookStore{DB: db}
}

var ErrNotEnoughStock = errors.New("not enough stock")

func (bs *BookStore) GetAllBooks() ([]Book, error) {
	rows, err := bs.DB.Query("SELECT id, title, author, copies FROM books")
	if err != nil {
		return nil, err
	}
	_ = rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Copies)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (bs *BookStore) GetBook(ID string) (Book, error) {
	var b Book
	err := bs.DB.QueryRow("SELECT id, title, author, copies FROM books WHERE id = ?", ID).
		Scan(&b.ID, &b.Title, &b.Author, &b.Copies)
	if err == sql.ErrNoRows {
		return b, fmt.Errorf("ID %q not found", ID)
	}
	return b, err
}

func (bs *BookStore) GetCopies(ID string) (int, error) {
	var copies int
	err := bs.DB.QueryRow("SELECT copies FROM books WHERE id = ?", ID).Scan(&copies)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("ID %q not found", ID)
	}
	return copies, err
}

func (bs *BookStore) AddCopies(ID string, copies int) (int, error) {
	tx, err := bs.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() { _ = tx.Rollback() }()

	var currentCopies int
	err = tx.QueryRow("SELECT copies FROM books WHERE id = ?", ID).Scan(&currentCopies)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("ID %q not found", ID)
	}
	if err != nil {
		return 0, err
	}

	newCopies := currentCopies + copies
	_, err = tx.Exec("UPDATE books SET copies = ? WHERE id = ?", newCopies, ID)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	return newCopies, err
}

func (bs *BookStore) SubCopies(ID string, copies int) (int, error) {
	tx, err := bs.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() { _ = tx.Rollback() }()

	var currentCopies int
	err = tx.QueryRow("SELECT copies FROM books WHERE id = ?", ID).Scan(&currentCopies)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("ID %q not found", ID)
	}
	if err != nil {
		return 0, err
	}

	if currentCopies < copies {
		return 0, fmt.Errorf("%w: %d", ErrNotEnoughStock, currentCopies)
	}

	newCopies := currentCopies - copies
	_, err = tx.Exec("UPDATE books SET copies = ? WHERE id = ?", newCopies, ID)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	return newCopies, err
}

func (bs *BookStore) AddBook(book Book) error {
	_, err := bs.DB.Exec("INSERT INTO books (id, title, author, copies) VALUES (?, ?, ?, ?)",
		book.ID, book.Title, book.Author, book.Copies)
	return err
}

func (bs *BookStore) SetCopies(ID string, copies int) error {
	if copies < 0 {
		return fmt.Errorf("negative number of copies: %d", copies)
	}
	_, err := bs.DB.Exec("UPDATE books SET copies = ? WHERE id = ?", copies, ID)
	return err
}
