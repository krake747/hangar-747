package handlers

import (
	"net/http"
	"strconv"

	"bookstore2/internal/models"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	BookStore *models.BookStore
}

func NewHandler(bs *models.BookStore) *Handler {
	return &Handler{BookStore: bs}
}

func (h *Handler) ListBooks(c echo.Context) error {
	books, err := h.BookStore.GetAllBooks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, books)
}

func (h *Handler) FindBook(c echo.Context) error {
	id := c.Param("id")
	book, err := h.BookStore.GetBook(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, book)
}

func (h *Handler) GetCopies(c echo.Context) error {
	id := c.Param("id")
	copies, err := h.BookStore.GetCopies(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, copies)
}

func (h *Handler) AddCopies(c echo.Context) error {
	id := c.Param("id")
	copiesStr := c.Param("copies")
	copies, err := strconv.Atoi(copiesStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid copies"})
	}
	newCopies, err := h.BookStore.AddCopies(id, copies)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, newCopies)
}

func (h *Handler) SubCopies(c echo.Context) error {
	id := c.Param("id")
	copiesStr := c.Param("copies")
	copies, err := strconv.Atoi(copiesStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid copies"})
	}
	newCopies, err := h.BookStore.SubCopies(id, copies)
	if err != nil {
		status := http.StatusNotFound
		if err == models.ErrNotEnoughStock {
			status = http.StatusBadRequest
		}
		return c.JSON(status, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, newCopies)
}
