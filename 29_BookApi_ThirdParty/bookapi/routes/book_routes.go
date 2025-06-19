package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jy-pandey/bookapi/handlers"
)

func RegisterRoutes(router *gin.Engine, h *handlers.BookHandler, topBookh *handlers.TopBooksHandler) {
	books := router.Group("/books")
	{
		books.GET("/", h.GetBooks)
		books.POST("/", h.CreateBook)
		books.GET("/:id", h.GetBookByID)
		books.PUT("/:id", h.UpdateBook)
		books.DELETE("/:id", h.DeleteBook)
	}

	top10books := router.Group("/top10books")
	{
		top10books.POST("/", topBookh.ImportTopTrendingBooks)
	}

	router.Static("/swagger", "./docs/swagger-ui")

	// serve yaml file
	// (route, filepath(from main.go)) 
	router.StaticFile("/docs/openapi.yaml", "./docs/openapi.yaml")
}