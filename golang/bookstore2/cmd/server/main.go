package main

import (
	"log"
	"os"

	"bookstore2/core"
	"bookstore2/core/books"

	"github.com/labstack/echo/v4"
)

func main() {
	dbPath := "books.db"
	if len(os.Args) > 1 {
		dbPath = os.Args[1]
	}

	db, err := core.NewDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bookStore := books.NewBookStore(db)
	h := books.NewHandler(bookStore)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome to the Bookstore API")
	})
	e.GET("/v1/list", h.ListBooks)
	e.GET("/v1/find/:id", h.FindBook)
	e.GET("/v1/getcopies/:id", h.GetCopies)
	e.GET("/v1/addcopies/:id/:copies", h.AddCopies)
	e.GET("/v1/subcopies/:id/:copies", h.SubCopies)

	addr := "localhost:3000"
	log.Printf("Starting server on %s with DB %s", addr, dbPath)
	log.Fatal(e.Start(addr))
}
