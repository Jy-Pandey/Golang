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
	// StuPerCourse(c)

	//* find grades of studnts in each subject
	// FindGrades(c)

	//* Preload
	// var students []models.Student
	// err := DB.Preload("Courses").Find(&students).Error
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, students)

	//* Easy level
	Easy(c)

}
func Easy(c *gin.Context) {
	// todo : Get the first 3 students sorted by age (youngest first).
	// var students []models.Student

	// err := DB.Find(&students).
	// Order("age asc").
	// Limit(3).Error

	// todo : Find students whose name contains the letter "a" (case-insensitive).
	// var students []models.Student
	// err := DB.Where("name ILIKE ?", "%i%").Find(&students).Error

	// todo : Get students whose age is either 20 or 22.
	// var students []models.Student
	// err := DB.Where("age IN ?", []int{20, 22}).Find(&students).Error
	// err := DB.Where("age = ? OR age = ?", 20, 22).Find(&students).Error

	// todo : Count how many students are enrolled in total.
	var enrolledStudents uint
	err := DB.Model(&models.StudentCourse{}).
			Select("count(*) as enrolledStudents").
			Scan(&enrolledStudents).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enrolledStudents)
}
func FindGrades(c *gin.Context) {
	type CourseGrade struct {
		Id      uint
		Name    string
		Subject string
		Grade   string
	}

	var result []CourseGrade

	err := DB.Model(&models.Student{}).
		Joins("Join student_courses AS sc on students.id = sc.student_id").
		Joins("Join courses AS c on sc.course_id = c.id").
		Select("students.id, students.name, c.title as subject, sc.grade").
		Scan(&result).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
func StuPerCourse(c *gin.Context) {

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
