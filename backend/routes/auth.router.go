package routes

import (
	"github.com/gin-gonic/gin"
	"backend/controllers"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/login", controllers.Login)
}
