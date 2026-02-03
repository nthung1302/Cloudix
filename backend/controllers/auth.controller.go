package controllers

import (
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := 1

	accessToken, _ := utils.GenerateAccessToken(userID, req.Username)

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
	})

	result, err := repositories.CreateUserByProc(req.Username, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if result.Result != 1 {
		c.JSON(400, gin.H{"message": result.Message})
		return
	}

	c.JSON(200, gin.H{"message": result.Message})
}
