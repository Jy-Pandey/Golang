package db

import (
	"fmt"
	"log"

	"github.com/jy-pandey/bookapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=jyoti dbname=bookapi port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatalf("Auto migration failed %v", err)
	}

	fmt.Println("Database connection established")
	return db
}
