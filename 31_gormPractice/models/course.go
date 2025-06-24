package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title       string
	Description string
	Students    []Student `gorm:"many2many:student_courses"`
}
