package main

import (
	"fmt"
	"gormPractice/handlers"
	"gormPractice/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Learn gorm raw-queries")
	dsn := "host=localhost user=postgres password=jyoti dbname=gormPractice port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.Student{}, &models.Course{}, &models.StudentCourse{})
	if err != nil {
		log.Fatalf("Auto migration failed %v", err)
	}

	fmt.Println("Database connection established")

	// student1 := models.Student{Name: "Jyoti", Email: "jyoti@gmail.com", Age: 20}
	// student2 := models.Student{Name: "Amit", Email: "amit@example.com", Age: 22}
	// student3 := models.Student{Name: "Ravi", Email: "ravi@example.com", Age: 21}

	// db.Create(&student1)
	// db.Create(&student2)
	// db.Create(&student3)

	// // Create courses
	// course1 := models.Course{Title: "Math", Description: "Algebra and Calculus"}
	// course2 := models.Course{Title: "Science", Description: "Physics and Chemistry"}
	// course3 := models.Course{Title: "History", Description: "World history"}

	// db.Create(&course1)
	// db.Create(&course2)
	// db.Create(&course3)

	// // Enroll students using join table
	// db.Create(&models.StudentCourse{StudentID: student1.ID, CourseID: course1.ID, EnrolledOn: time.Now(), Grade: "A"})
	// db.Create(&models.StudentCourse{StudentID: student1.ID, CourseID: course2.ID, EnrolledOn: time.Now(), Grade: "B"})
	// db.Create(&models.StudentCourse{StudentID: student2.ID, CourseID: course1.ID, EnrolledOn: time.Now(), Grade: "C"})
	// db.Create(&models.StudentCourse{StudentID: student3.ID, CourseID: course3.ID, EnrolledOn: time.Now(), Grade: "A"})

	
	r := gin.Default()

	handlers.IntializeDB(db)
	r.POST("/gorm", handlers.PracticeGorm)

	r.Run(":8080")
}
