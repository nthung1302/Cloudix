package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func AccountRoutes(r *gin.RouterGroup) {
	r.POST("/change-password", controllers.ChangePassword)
}
