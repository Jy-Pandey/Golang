package main

import (
	"bookstore/db"
	"bookstore/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	// db.InitDB() // DB ko initialize kiya (global DB ban gaya)

	r := gin.Default()

	r.GET("/books", handlers.GetBooks)
	r.POST("/books", handlers.CreateBook)

	r.Run() // start server
}