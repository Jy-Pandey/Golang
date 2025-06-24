package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name    string
	Email   string `gorm:"unique"`
	Age     int
	Courses []Course `gorm:"many2many:student_courses"`
}
