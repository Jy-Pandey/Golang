package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Book struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
}

var books = []Book{
    {ID: 1, Title: "Book One"},
    {ID: 2, Title: "Book Two"},
}

func main() {
    r := gin.Default()

    r.GET("/books", func(c *gin.Context) {
        c.JSON(http.StatusOK, books)
    })

    // Serve OpenAPI YAML file
    r.StaticFile("/openapi.yaml", "./openapi.yaml")

    // Serve Swagger UI files
    r.Static("/docs", "./static/swagger-ui")

    r.Run(":8080")
}
