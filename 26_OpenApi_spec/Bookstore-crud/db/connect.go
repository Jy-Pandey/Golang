package db

import (
	"bookstore-crud/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=jyoti dbname=bookstore port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect DB : ", err) // stop and exit
	}

	db.AutoMigrate(&models.Book{})
	return db
}
