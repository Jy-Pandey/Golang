package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model                  // provide extra fields - id, created/updated/deleted At
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Students    []StudentCourse `json:"students,omitempty"`
}
