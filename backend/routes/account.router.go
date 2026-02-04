package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func AccountRoutes(r *gin.RouterGroup) {
	r.GET("/profile", controllers.GetProfile)
	r.POST("/change-password", controllers.ChangePassword)
}
