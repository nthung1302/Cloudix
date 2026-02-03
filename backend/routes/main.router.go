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
	}
}