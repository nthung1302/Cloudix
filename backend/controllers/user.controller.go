package controllers

import (

	"github.com/gin-gonic/gin"
)

func ChangePassword(c *gin.Context) {
	
}

func GetProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(int)
	username := c.MustGet("username").(string)

	c.JSON(200, gin.H{
		"user_id":  userID,
		"username": username,
	})
}