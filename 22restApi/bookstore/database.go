package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Db is a pointer to gorm.DB object //! Global
// It is used in whole application to interact with database like - insert, read, update, delete etc.

func InitDB() {
	// DSN (Data Source Name)
	dsn := "host=localhost user=postgres password=jyoti dbname=bookstore port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database : ", err) // immediately exit the program
	}
}
