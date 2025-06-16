package handlers

import (
	"looselyCoupled/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookHandler struct {
	DB *gorm.DB // Injected dependency
}

// Constructor function – dependency injected here
func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{DB: db}
}

// GET /books
func (h *BookHandler) GetBooks(c *gin.Context) { // gin.Context - handle request response
	var books []models.Book
	h.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

// POST /books
func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.DB.Create(&book)
	c.JSON(http.StatusOK, book)
}
