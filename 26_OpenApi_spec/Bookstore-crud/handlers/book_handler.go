package handlers

import (
	"bookstore-crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookHandler struct {
	DB *gorm.DB
}

// NewHandler returns a handler with injected DB
func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{DB: db}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	var books []models.Book
	h.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	var book models.Book
	idParam := c.Param("id")

    id, err := strconv.Atoi(idParam) // string → int
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
	if err := h.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}
func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}
func (h *BookHandler) UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := h.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.Title = input.Title
	book.Author = input.Author
	h.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}
func (h *BookHandler) DeleteBook(c *gin.Context) {
	var book models.Book
	if err := h.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	h.DB.Delete(&book)
	c.Status(http.StatusNoContent)
}
