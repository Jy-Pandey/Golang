package models

import "time"

type StudentCourse struct {
	StudentId  uint `json:"student_id"`
	CourseId   uint `json:"course_id"`
	EnrolledOn time.Time `json:"enrolled_on"`
}