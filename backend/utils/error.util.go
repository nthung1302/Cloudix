// utils/error.go
package utils

import (
	"os"
	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, httpCode int, message string, err error) {
	if os.Getenv("DEBUG") == "true" {
		c.JSON(httpCode, gin.H{
			"message": message,
			"debug":   err.Error(),
		})
		return
	}

	c.JSON(httpCode, gin.H{
		"message": message,
	})
}
