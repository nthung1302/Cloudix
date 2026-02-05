package routes

import (
	"backend/middlewares"

	"github.com/gin-gonic/gin"
)

func MainRoutes(r *gin.Engine) {
	api := r.Group("/api/v1.0")
	{
		auth := api.Group("/authentication")
		{
			AuthRoutes(auth)
		}

		account := api.Group("/account", middlewares.JWTAuth())
		{
			AccountRoutes(account)
		}
	}
}
