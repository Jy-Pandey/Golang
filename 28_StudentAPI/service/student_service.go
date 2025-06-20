package service

import (
	"errors"
	"studentApi/models"
	"studentApi/repository"
)

type StudentService struct {
	StudentRepo *repository.StudentRepository
	CourseRepo *repository.CourseRepository
}

func NewStudentService(sr *repository.StudentRepository, cr *repository.CourseRepository) *StudentService{
	return &StudentService{StudentRepo: sr, CourseRepo: cr}
}

// Add new Student
func (s *StudentService) CreateStudent(student *models.Student) error {
	return s.StudentRepo.CreateStudent(student)
}

// Get all students
func (s *StudentService) GetAllStudents() ([]models.Student, error) {
	return s.StudentRepo.GetAllStudents()
}

// Enroll a student in a course
func(s *StudentService) EnrollStudent(studentID, courseID uint) error {

	// Step 1 : check if student exist
	studentExist := s.StudentRepo.CheckStudentExists(studentID) 
	if !studentExist {
		return errors.New("student does not exist")
	}

	// Check course exist
	courseExist := s.CourseRepo.CheckCourseExists(courseID)
	if !courseExist {
		return errors.New("course does not exist")
	}

	// Step 3 : Enroll student
	return s.StudentRepo.EnrollStudentInCourse(studentID, courseID)
}

// Get all courses of a studnt
func (s *StudentService) GetStudentCourses(studentID uint) ([]models.Course, error) {
	return s.StudentRepo.GetStudentCourses(studentID)
}