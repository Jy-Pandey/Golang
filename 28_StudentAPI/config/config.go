package config

import (
	"log"
	"studentApi/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=jyoti dbname=student_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}
	return db
}
func MigrateTables(db *gorm.DB) {
	db.AutoMigrate(&models.Student{}, &models.Course{}, &models.StudentCourse{}, &models.User{})
}
