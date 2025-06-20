package repository

import (
	"studentApi/models"

	"gorm.io/gorm"
)

type CourseRepository struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{DB: db}
}

// Create new Course
func (r *CourseRepository) CreateCourse(course *models.Course) error {
	return r.DB.Create(course).Error
}
func (r *CourseRepository) CheckCourseExists(courseID uint) bool {
	var course models.Course
	err := r.DB.First(&course, courseID).Error
	return err == nil
}

// Get all courses
func (r *CourseRepository) GetAllCourses() ([]models.Course, error) {
	var courses []models.Course
	err := r.DB.Find(&courses).Error
	return courses, err

}
