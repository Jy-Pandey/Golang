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
    {ID: 3, Title: "Book three"},
    {ID: 4, Title: "Book Four"},
}

func main() {
    r := gin.Default()

    r.GET("/books", func(c *gin.Context) {
        c.JSON(http.StatusOK, books)
    })

    // Serve OpenAPI YAML file
    // Used to serve a single file 
    // When user visits http://localhost:8080/openapi.yaml Gin sends the file openapi.yaml
    r.StaticFile("route", "path")
    r.StaticFile("/openapi.yaml", "./openapi.yaml")

    // Serve Swagger UI files
    // Used to serve all files inside a folder
    // When user visits http://localhost:8080/docs
    // It will load index.html from inside static/swagger-ui
    // And any other files (CSS, JS) are also served automatically
    r.Static("/docs", "./static/swagger-ui")

    r.Run(":8080")
}
