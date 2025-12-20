package main

import (
	"log/slog"
	"net/http"
	"os"

	"bookstore2/core"
	"bookstore2/core/books"
	"bookstore2/core/shared"

	"github.com/labstack/echo/v4"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	dbPath := "books.db"
	if len(os.Args) > 1 {
		dbPath = os.Args[1]
	}

	logger.Info("Initializing database", "path", dbPath)
	db, err := core.NewDB(dbPath)
	if err != nil {
		logger.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}
	defer func() { _ = db.Close() }()

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
	if err := e.Start(addr); err != nil {
		logger.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
