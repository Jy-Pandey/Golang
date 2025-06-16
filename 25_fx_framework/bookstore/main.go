package main

import (
	"bookstore/db"
	"bookstore/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// This function creates a new Gin router (type *gin.Engine)
// It defines the API routes and connects them to handler functions
func NewRouter(handler *handlers.BookHandler) *gin.Engine {
	r := gin.Default() // Creates a Gin engine with Logger and Recovery middleware

	// Define routes
	r.GET("/books", handler.GetBooks)    // GET request → list of books
	r.POST("/books", handler.CreateBook) // POST request → create new book

	return r // Return the configured Gin engine
}

// This is the entry point of the application using the Fx framework.
// Fx helps us manage and inject dependencies automatically (like DB, handlers, router).
func main() {
	fmt.Println("Learn FX framework")

	app := fx.New(
		fx.Provide(
			db.InitDB,               // Provides *gorm.DB (a PostgreSQL connection)
			handlers.NewBookHandler, // Provides *BookHandler (requires *gorm.DB)
			NewRouter,               // Provides *gin.Engine (requires *BookHandler)
		),
		fx.Invoke(
			// Starts the Gin server after everything is ready
			func(r *gin.Engine) {
				r.Run(":8080") // Run the server on port 8080
			},
		),
	)

	app.Run() // Start the Fx application (runs all providers and invokes)
}
