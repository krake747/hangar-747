package core

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"bookstore2/core/books"
	"bookstore2/core/shared"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func RunServer(dbPath string) error {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("Initializing database", "path", dbPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer func() { _ = db.Close() }()

	initSQL, err := os.ReadFile("testdata/init.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(initSQL))
	if err != nil {
		return err
	}

	bookStore := books.NewBookStore(db)
	h := books.NewHandler(logger, bookStore)

	e := echo.New()

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if pd, ok := err.(*shared.ProblemDetail); ok {
			_ = c.JSON(pd.Status, pd)
		} else {
			_ = c.JSON(
				http.StatusInternalServerError,
				shared.NewProblemDetail(http.StatusInternalServerError, "Internal Server Error", err.Error()),
			)
		}
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome to the Bookstore API")
	})
	e.GET("/v1/list", h.ListBooks)
	e.GET("/v1/find/:id", h.FindBook)
	e.GET("/v1/getcopies/:id", h.GetCopies)
	e.POST("/v1/addcopies/:id/:copies", h.AddCopies)
	e.POST("/v1/subcopies/:id/:copies", h.SubCopies)

	addr := "localhost:3000"
	logger.Info("Starting server", "addr", addr, "dbPath", dbPath)
	logger.Info("Server started successfully", "addr", addr)
	return e.Start(addr)
}
