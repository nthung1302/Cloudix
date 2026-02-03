package routes
import (
	
	"github.com/gin-gonic/gin"
)

func MainRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth := api.Group("/authentication")
		{
			AuthRoutes(auth)
		}

		/*
		user := api.Group("/users")
		{
			UserRoutes(middlewares.JWTAuth, user)
		}

		product := api.Group("/products")
		{
			ProductRoutes(middlewares.JWTAuth, product)
		}
			*/
	}
}