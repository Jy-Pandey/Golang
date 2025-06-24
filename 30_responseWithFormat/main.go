package main

import (
	"fmt"
	"formatResponse/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func main() {
	fmt.Println("Lets Learn how to send response in format")
	r := gin.Default()

	r.POST("/students", func(c *gin.Context) {
		var student Student

		if err := c.ShouldBindJSON(&student); err != nil {
			utils.RespondError(c, http.StatusBadRequest, "Invalid request", err.Error())
			return
		}

		// perform operation if any
		utils.RespondSuccess(c, http.StatusOK, "Student is added", student)
	})

	// Api to store file to our sever
	r.POST("/upload", func(c *gin.Context) {
		// 1. Read the file
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": "No file uploaded"})
			return
		}

		// 2. save the files to uploads/ folder
		err = c.SaveUploadedFile(file, "./uploads/"+file.Filename)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save file"})
			return
		}

		// 3. Return the filepath
		utils.RespondSuccess(c, http.StatusAccepted, "File uploaded successfully", map[string]string{
			"url": "http://localhost:8080/images/" + file.Filename,
		})
	})

	// http://localhost:8080/images/wallpaper.webp
	r.Static("/images", "./uploads")
	r.Run(":8080")
}
