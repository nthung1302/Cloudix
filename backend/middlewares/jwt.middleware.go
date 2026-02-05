package middlewares

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
			utils.Logger.Error("Không tìm thấy token trong header.")
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
			utils.Logger.Error("Token không hợp lệ: ", err)
			return
		}

		// 3. Lấy thông tin user từ token
		claims := token.Claims.(*utils.Claims)
		c.Set("UID", claims.UID)
		c.Set("UserName", claims.UserName)
		c.Next()
	}
}
