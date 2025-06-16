package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Reast Api using Gin framework")

	r := gin.Default()
	InitDB()
	DB.AutoMigrate(&Book{}) // create table automatically

	r.GET("/books", GetBooks)
	r.POST("/books", CreateBook)
	r.PUT("/books/:id", UpdateBook)
    r.DELETE("/books/:id", DeleteBook)

    r.Run(":8080") // http://localhost:8080
}
