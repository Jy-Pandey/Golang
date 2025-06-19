package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model                 // provide extra fields - id, created/updated/deleted At
	Name       string          `json:"name"`
	Age        int             `json:"age"`
	Email      string          `json:"email"`
	Courses    []StudentCourse `json:"courses,omitempty"`
}
