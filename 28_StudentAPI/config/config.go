package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=jyoti dbname=student_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if(err != nil) {
		log.Fatal("Failed to connect database")
	}
	return db
}