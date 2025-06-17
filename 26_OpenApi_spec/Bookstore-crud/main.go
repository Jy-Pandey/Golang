package main

import (
	"bookstore-crud/db"
	"bookstore-crud/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.ConnectDB()
	h := handlers.NewBookHandler(database) // Inject DB once

	r := gin.Default()

	// Routes
	r.GET("/books", h.GetBooks)
	r.GET("books/:id", h.GetBookByID)
	r.POST("/books", h.CreateBook)
	r.PUT("/books/:id", h.UpdateBook)
	r.DELETE("/books/:id", h.DeleteBook)


	// (route, filepath)
	r.Static("/docs", "./static/swagger-ui")
	r.StaticFile("/openapi.yaml", "./openapi.yaml")
	
	r.Run(":8080")
}
