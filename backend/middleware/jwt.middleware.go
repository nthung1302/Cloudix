package middleware

import (
	"backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Lấy token từ header
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Thiếu token"})
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		// 2. Verify token
		token, err := jwt.ParseWithClaims(
			tokenStr,
			&utils.Claims{},
			func(t *jwt.Token) (interface{}, error) {
				return utils.AccessSecret, nil
			},
		)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token không hợp lệ hoặc hết hạn"})
			return
		}

		// 3. Lấy thông tin user từ token
		claims := token.Claims.(*utils.Claims)
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}

func RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.ParseWithClaims(
		req.RefreshToken,
		&utils.Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return utils.RefreshSecret, nil
		},
	)

	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Refresh token không hợp lệ"})
		return
	}

	claims := token.Claims.(*utils.Claims)
	newAccess, _ := utils.GenerateAccessToken(claims.UserID, claims.Username)

	c.JSON(200, gin.H{
		"access_token": newAccess,
	})
}
