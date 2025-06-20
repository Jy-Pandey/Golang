package handler

import (
	"net/http"
	"studentApi/models"
	"studentApi/service"

	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	Service *service.CourseService
}

func NewCourseHandler(s *service.CourseService) *CourseHandler {
	return &CourseHandler{Service: s}
}

// POST /courses
func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var course models.Course

	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateCourse(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create course"})
		return
	}

	c.JSON(http.StatusCreated, course)
}
// GET /courses
func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	courses, err := h.Service.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}
