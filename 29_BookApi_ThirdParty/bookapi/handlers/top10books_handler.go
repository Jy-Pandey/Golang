package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jy-pandey/bookapi/models"
	"gorm.io/gorm"
)

type TopBooksHandler struct {
	DB *gorm.DB
}

func NewTopBooksHandler(db *gorm.DB) *TopBooksHandler {
	return &TopBooksHandler{DB: db}
}

type DummyProduct struct {
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
}

// Schema for DummyProducts
type DummyResponse struct {
	Products []DummyProduct `json:"products"`
	Total    int               `json:"total"`
}

func (h *TopBooksHandler) ImportTopTrendingBooks(c *gin.Context) {

	res, err := http.Get("https://dummyjson.com/products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch trending books"})
		return
	}
	defer res.Body.Close()

	var response DummyResponse // to store API data
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response format"})
        return
	}

	books := make([]models.Book, 0)

	// store top 10 products data into books then save to DB
	for i, product := range response.Products {
		if i > 10 {
			break
		}
		books = append(books, models.Book{
			Title: product.Title,
			Description: product.Description,
			Price: product.Price,
			Author: "Jyoti Pandey",
		})
	}

	if err := h.DB.Create(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Top 10 trending books imported", "count": len(books)})
}
