package utils

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func DebugLogErr(c *gin.Context, status int, error_code int, message error) {
	if os.Getenv("DEBUG") == "true" {
		c.JSON(status, gin.H{
			"status":  status,
			"error":   error_code,
			"message": message,
			"debug":   message.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"status":  status,
		"error":   error_code,
		"message": message,
	})
}

func DebugLog(c *gin.Context, status int, error_code int, message string) {
	c.JSON(status, gin.H{
		"status":  status,
		"error":   error_code,
		"message": message,
	})
}

func DebugTerminal(c *gin.Context, title, message string) {
	log.Printf("%s: %s", title, message+" | "+time.Now().Format(time.RFC3339))
}
