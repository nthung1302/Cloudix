package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/login", controllers.LoginAccount)
	r.POST("/register", controllers.RegisterAccount)
}
