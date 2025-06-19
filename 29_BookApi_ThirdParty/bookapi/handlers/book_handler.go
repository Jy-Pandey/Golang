package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jy-pandey/bookapi/models"
	"gorm.io/gorm"
)

type BookHandler struct {
	DB *gorm.DB
}

func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{DB : db}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	var books []models.Book
	if err := h.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
	}
	c.JSON(http.StatusCreated, book)
}
func (h *BookHandler) GetBookByID(c *gin.Context) {
    id := c.Param("id")
    var book models.Book

    if err := h.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    c.JSON(http.StatusOK, book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book

    if err := h.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    var updatedBook models.Book
    if err := c.ShouldBindJSON(&updatedBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    book.Title = updatedBook.Title
    book.Author = updatedBook.Author
    book.Description = updatedBook.Description
    book.Price = updatedBook.Price

    if err := h.DB.Save(&book).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
    id := c.Param("id")
    if err := h.DB.Delete(&models.Book{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
