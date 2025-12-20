package books

import (
	"log/slog"
	"net/http"
	"strconv"

	"bookstore2/core/shared"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Logger    *slog.Logger
	BookStore *BookStore
}

func NewHandler(logger *slog.Logger, bs *BookStore) *Handler {
	return &Handler{
		Logger:    logger,
		BookStore: bs,
	}
}

func (h *Handler) ListBooks(c echo.Context) error {
	h.Logger.Info("Listing all books")
	books, err := h.BookStore.GetAllBooks()
	if err != nil {
		h.Logger.Error("Failed to list books", "error", err)
		c.Error(shared.NewProblemDetail(http.StatusInternalServerError, "Internal Server Error", err.Error()))
		return nil
	}
	h.Logger.Info("Listed books successfully", "count", len(books))
	return c.JSON(http.StatusOK, books)
}

func (h *Handler) FindBook(c echo.Context) error {
	id := c.Param("id")
	h.Logger.Info("Finding book", "id", id)
	book, err := h.BookStore.GetBook(id)
	if err != nil {
		h.Logger.Error("Failed to find book", "id", id, "error", err)
		c.Error(shared.NewProblemDetail(http.StatusNotFound, "Book Not Found", err.Error()))
		return nil
	}
	h.Logger.Info("Found book", "id", id, "title", book.Title)
	return c.JSON(http.StatusOK, book)
}

func (h *Handler) GetCopies(c echo.Context) error {
	id := c.Param("id")
	h.Logger.Info("Getting copies for book", "id", id)
	copies, err := h.BookStore.GetCopies(id)
	if err != nil {
		h.Logger.Error("Failed to get copies", "id", id, "error", err)
		c.Error(shared.NewProblemDetail(http.StatusNotFound, "Book Not Found", err.Error()))
		return nil
	}
	h.Logger.Info("Retrieved copies", "id", id, "copies", copies)
	return c.JSON(http.StatusOK, copies)
}

func (h *Handler) AddCopies(c echo.Context) error {
	id := c.Param("id")
	copiesStr := c.Param("copies")
	copies, err := strconv.Atoi(copiesStr)
	if err != nil {
		h.Logger.Error("Invalid copies parameter", "copiesStr", copiesStr, "error", err)
		c.Error(shared.NewProblemDetail(http.StatusBadRequest, "Bad Request", "Invalid copies parameter"))
		return nil
	}
	h.Logger.Info("Adding copies to book", "id", id, "copies", copies)
	newCopies, err := h.BookStore.AddCopies(id, copies)
	if err != nil {
		h.Logger.Error("Failed to add copies", "id", id, "copies", copies, "error", err)
		c.Error(shared.NewProblemDetail(http.StatusNotFound, "Book Not Found", err.Error()))
		return nil
	}
	h.Logger.Info("Added copies successfully", "id", id, "newCopies", newCopies)
	return c.JSON(http.StatusOK, newCopies)
}

func (h *Handler) SubCopies(c echo.Context) error {
	id := c.Param("id")
	copiesStr := c.Param("copies")
	copies, err := strconv.Atoi(copiesStr)
	if err != nil {
		h.Logger.Error("Invalid copies parameter", "copiesStr", copiesStr, "error", err)
		c.Error(shared.NewProblemDetail(http.StatusBadRequest, "Bad Request", "Invalid copies parameter"))
		return nil
	}
	h.Logger.Info("Subtracting copies from book", "id", id, "copies", copies)
	newCopies, err := h.BookStore.SubCopies(id, copies)
	if err != nil {
		if err == ErrNotEnoughStock {
			h.Logger.Warn("Not enough stock", "id", id, "requested", copies)
			c.Error(shared.NewProblemDetail(http.StatusBadRequest, "Insufficient Stock", err.Error()))
		} else {
			h.Logger.Error("Failed to subtract copies", "id", id, "copies", copies, "error", err)
			c.Error(shared.NewProblemDetail(http.StatusNotFound, "Book Not Found", err.Error()))
		}
		return nil
	}
	h.Logger.Info("Subtracted copies successfully", "id", id, "newCopies", newCopies)
	return c.JSON(http.StatusOK, newCopies)
}
