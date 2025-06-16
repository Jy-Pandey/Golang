package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
// | Purpose                 | Code Example                                  |
// | ----------------------- | --------------------------------------------- |
// | Request parameters lena | `c.Param("id")`                               |
// | Query string lena       | `c.Query("name")`                             |
// | JSON body read karna    | `c.BindJSON(&yourStruct)`                     |
// | Response bhejna         | `c.JSON(...)`, `c.String(...)`, `c.HTML(...)` |
// | Status code set karna   | `c.Status(http.StatusBadRequest)`             |


// | Data Source  | Method                           | Content-Type                        |
// | ------------ | -------------------------------- | ----------------------------------- |
// | JSON body    | `ShouldBindJSON()`               | `application/json`                  |
// | Form fields  | `ShouldBind()`                   | `application/x-www-form-urlencoded` |
// | Query string | `ShouldBindQuery()`              | In URL                              |
// | Path params  | `ShouldBindUri()`                | URL path params like `/user/:id`    |
// | File Upload  | `FormFile()` / `MultipartForm()` | `multipart/form-data`               |

// ShouldBind vs Bind	ShouldBind aapko error handle karne ka control deta hai, Bind automatic 400 error deta hai

func GetBooks(c *gin.Context) { // this c control the functionality of request/response
	var books []Book // slice of Book struct
	DB.Find(&books) // go find and store in books
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	}
	DB.Create(&book)
	c.JSON(http.StatusOK, book)

}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book Book
	if err := DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
	}

	// Take updated data
	var input Book
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	book.Title = input.Title
    book.Author = input.Author
    DB.Save(&book)

    c.JSON(http.StatusOK, book)

}
func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    var book Book
    if err := DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    DB.Delete(&book)
    c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}