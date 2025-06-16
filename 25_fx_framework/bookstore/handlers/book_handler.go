package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookHandler struct {
	db *gorm.DB
}

func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{db: db}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	var books []Book
	h.db.Find(&books)
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.db.Create(&book)
	c.JSON(http.StatusCreated, book)
}

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
