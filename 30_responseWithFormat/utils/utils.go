package utils

import "github.com/gin-gonic/gin"

func RespondSuccess(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"success" : true,
		"message" : message,
		"data" : data,
	})
}
func RespondError(c *gin.Context, status int, message string, err interface{}) {
	c.JSON(status, gin.H{
		"success" : false,
		"message" : message,
		"data" : err,
	})
}