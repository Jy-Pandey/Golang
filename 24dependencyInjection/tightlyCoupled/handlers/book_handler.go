package handlers

import (
	"bookstore/db"
	"bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
    var books []models.Book
    db.DB.Find(&books) // ❗ Global DB used
    c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
    var book models.Book

    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.DB.Create(&book) // ❗ Global DB used
    c.JSON(http.StatusCreated, book)
}
