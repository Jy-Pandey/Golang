package routes

import (
	"studentApi/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, studentHandler *handler.StudentHandler, courseHandler *handler.CourseHandler) {
	studentRoutes := router.Group("/students")
	{
		studentRoutes.GET("", studentHandler.GetAllStudents)
		studentRoutes.POST("", studentHandler.CreateStudent)
		studentRoutes.POST("/:studentId/enroll/:courseId", studentHandler.EnrollStudent)
		studentRoutes.GET("/:id/courses", studentHandler.GetStudentCourses)
	}

	courseRoutes := router.Group("/courses")
	{
		courseRoutes.GET("", courseHandler.GetAllCourses)
		courseRoutes.POST("", courseHandler.CreateCourse)
	}

	// route , filepath
	router.Static("/swagger", "./docs/swagger-ui")

	router.StaticFile("/docs/openapi.yaml", "./docs/openapi.yaml")
}
