package main

import (
	"fmt"

	"studentApi/auth"
	"studentApi/config"
	"studentApi/handler"
	"studentApi/repository"
	"studentApi/routes"
	"studentApi/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	fmt.Println("Welcome to Student API")

	app := fx.New(
		// provide dependencies
		fx.Provide(
			config.ConnectDatabase,          //* provides *gorm.DB
			repository.NewStudentRepository, // *StudentRepository
			repository.NewCourseRepository,  // *CourseRepository
			service.NewStudentService,       // *StudentService
			service.NewCourseService,       // *CourseService
			handler.NewStudentHandler,       // *StudentHandler
			handler.NewCourseHandler,        // *CourseHandler
			auth.NewAuthService,
			auth.NewAuthHandler,
			gin.Default,                     // *gin.Engine
		),
		fx.Invoke(
			config.MigrateTables,  // need db, take from provided dependency
			routes.RegisterRoutes, // take *router, studentHandler, courseHandler from dependency
			StartServer,
		),
	)

	app.Run()
}
func StartServer(router *gin.Engine) {
	router.Run(":8080") // https://localhost:8080/
}

// App startup
//    ↓
// Fx creates gorm.DB using config.ConnectDatabase
//    ↓
// Fx uses DB to create Repos
//    ↓
// Fx uses Repos to create Services
//    ↓
// Fx uses Services to create Handlers
//    ↓
// Fx calls RegisterRoutes() via fx.Invoke and injects Router + Handlers
//    ↓
// Gin starts automatically
