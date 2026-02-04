package routes

import (
	"backend/middlewares"
	"github.com/gin-gonic/gin"
)

func MainRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth := api.Group("/authentication")
		{
			AuthRoutes(auth)
		}

		
		account := api.Group("/users", middlewares.JWTAuth())
		{
			AccountRoutes(account)
		}
/*
		product := api.Group("/products")
		{
			ProductRoutes(middlewares.JWTAuth, product)
		}
			*/
	}
}