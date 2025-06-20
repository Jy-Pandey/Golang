package service

import (
	"studentApi/models"
	"studentApi/repository"
)

type CourseService struct {
	CourseRepo *repository.CourseRepository
}

func NewCourseService(cr *repository.CourseRepository) *CourseService {
	return &CourseService{CourseRepo: cr}
}

// Add new Course
func (s *CourseService) CreateCourse(course *models.Course) error {
	return s.CourseRepo.CreateCourse(course)
}

// Get All courses
func (s *CourseService) GetAllCourses() ([]models.Course, error) {
	return s.CourseRepo.GetAllCourses()
}