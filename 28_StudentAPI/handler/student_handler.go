package handler

import (
	"net/http"
	"strconv"
	"studentApi/models"
	"studentApi/service"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	Service *service.StudentService
}

func NewStudentHandler(s *service.StudentService) *StudentHandler{
	return &StudentHandler{Service: s}
}

// post : /students
func (h *StudentHandler) CreateStudent(c *gin.Context) {
	// extract student detail from body
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	if err := h.Service.CreateStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create student"})
		return
	}

	c.JSON(http.StatusCreated, student)
}

// GET /students
func (h *StudentHandler) GetAllStudents(c *gin.Context) {
	students, err := h.Service.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch students"})
		return
	}

	c.JSON(http.StatusOK, students)
}

// POST /students/:studentId/enroll/:courseId
func (h *StudentHandler) EnrollStudent(c *gin.Context) {
	studentID , err1 := strconv.Atoi(c.Param("studentId"))
	courseID, err2 := strconv.Atoi(c.Param("courseId"))

	
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid IDs"})
		return
	}

	err := h.Service.EnrollStudent(uint(studentID), uint(courseID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "student enrolled successfully"})
}

// GET /students/:id/courses
func (h *StudentHandler) GetStudentCourses(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	courses, err := h.Service.GetStudentCourses(uint(studentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}