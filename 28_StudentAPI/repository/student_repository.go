package repository

import (
	"studentApi/models"
	"time"

	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

// constructor
func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{DB: db}
}

// Create a new student
// Send student struct that need to add in database
// Through StudentRepository instance r will be able to access the db
func (r *StudentRepository) CreateStudent(student *models.Student) error {
	return r.DB.Create(student).Error
}

// Get All Students
func (r *StudentRepository) GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	err := r.DB.Find(&students).Error
	return students, err
}

func (r *StudentRepository) CheckStudentExists(studentID uint) bool {
	var student models.Student
	err := r.DB.First(&student, studentID).Error
	return err == nil
}

// Enroll student in course
func (r *StudentRepository) EnrollStudentInCourse(studentId, courseId uint) error {
	enrollment := models.StudentCourse{
		StudentId:  studentId,
		CourseId:   courseId,
		EnrolledOn: time.Now(),
	}
	return r.DB.Create(enrollment).Error
}

// Get courses for a student
func (r *StudentRepository) GetStudentCourses(studentId uint) ([]models.Course, error) {
	var courses []models.Course
	err := r.DB.Table("students").
		Joins("JOIN student_courses ON students.id = student_courses.student_id").
		Joins("JOIN courses ON student_courses.course_id = courses.id").
		Where("students.id = ?", studentId).
		Select("courses.*").
		Scan(&courses).Error
	return courses, err
}
