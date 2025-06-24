package handlers

import (
	"fmt"
	"gormPractice/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func IntializeDB(db *gorm.DB) {
	DB = db
}

func PracticeGorm(c *gin.Context) {
	fmt.Println("Inside handler")

	//* Get all students
	// GetAllStu(c)

	//* Get student by email
	// GetStuByEmail(c)

	//* List students older than 20
	// StuOlder(c)

	//* Find students enrolled in Math
	// StuEnrollMaths(c)

	//* Find courses taken by Jyoti
	// CoursesByStu(c)

	//* Count number of students per course
	StuPerCourse(c)

}
func StuPerCourse(c *gin.Context) {
	// select c.title, count(*)
	// from courses c
	// join student_courses sc on c.id = sc.course_id
	// group by c.title;
	type CourseCnt struct {
		Title string
		Total uint
	}
	var result []CourseCnt
	err := DB.Table("courses").
		Joins("Join student_courses AS sc on courses.id = sc.course_id").
		Group("courses.title").
		Select("courses.title, count(*) as total").
		Scan(&result).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
func CoursesByStu(c *gin.Context) {
	var courses []models.Course
	err := DB.
		Joins("Join student_courses AS sc on courses.id = sc.course_id").
		Joins("Join students AS s on s.id = sc.student_id").
		Where("s.id = ?", 1).
		Find(&courses).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, courses)
}
func StuEnrollMaths(c *gin.Context) {
	var students []models.Student
	err := DB.
		Joins("Join student_courses AS sc on students.id = sc.student_id").
		Joins("Join courses AS c on c.id = sc.course_id").
		Where("c.Title = ?", "Math").
		Find(&students).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}
func StuOlder(c *gin.Context) {
	var students []models.Student
	err := DB.Where("Age > ?", 20).Find(&students).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}
func GetStuByEmail(c *gin.Context) {
	var student models.Student
	err := DB.Where("Email = ?", "jyoti@gmail.com").First(&student).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}
func GetAllStu(c *gin.Context) {
	var students []models.Student
	err := DB.Find(&students).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}
