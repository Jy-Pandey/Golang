package main

import (
	"looselyCoupled/db"
	"looselyCoupled/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // return a gin.Engine()( => router object ) that define routes - GET, POST etc 
	// performs : Routing, middleware manage, server run

	// Step 1: DB create (not global)
	// Now if we want to test handlers we can provide a fake or dummy dataBase from here no need to change actual databse and handlers
	database := db.InitDB()

	// Step 2: Handler me DB inject
	bookHandler := handlers.NewBookHandler(database)

	// Step 3: Routes
	r.GET("/books", bookHandler.GetBooks)
	r.POST("/books", bookHandler.CreateBook)

	// Start server
	r.Run(":8080")
}
