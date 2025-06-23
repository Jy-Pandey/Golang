package routes

import (
	"studentApi/auth"
	"studentApi/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, studentHandler *handler.StudentHandler, courseHandler *handler.CourseHandler, authHandler *auth.AuthHandler) {
	// Public auth routes (no middleware)
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/signup", authHandler.Signup)
		authRoutes.POST("/login", authHandler.Login)
	}
	
	// Protected routes group - apply middleware here
	protected := router.Group("/")
	protected.Use(auth.AuthMiddleware()) // Applied to all inner routes
	{
		studentRoutes := protected.Group("/students")
		{
			studentRoutes.GET("", studentHandler.GetAllStudents)
			studentRoutes.POST("", studentHandler.CreateStudent)
			studentRoutes.POST("/:studentId/enroll/:courseId", studentHandler.EnrollStudent)
			studentRoutes.GET("/:id/courses", studentHandler.GetStudentCourses)
		}

		courseRoutes := protected.Group("/courses")
		{
			courseRoutes.GET("", courseHandler.GetAllCourses)
			courseRoutes.POST("", courseHandler.CreateCourse)
		}
	}
	

	// route , filepath
	router.Static("/swagger", "./docs/swagger-ui")

	router.StaticFile("/docs/openapi.yaml", "./docs/openapi.yaml")
}
