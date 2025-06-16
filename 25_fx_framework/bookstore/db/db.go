package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=jyoti dbname=bookstore port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connect failed:", err)
	}
	return db
}
