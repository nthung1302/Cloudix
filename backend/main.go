package main

import (
	"backend/configs"
	"backend/routes"
	"backend/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadEnv()
	configs.ConnectDB()
	utils.InitLogger()

	r := gin.Default()

	routes.MainRoutes(r)

	if err := r.Run(":" + os.Getenv("SERVER_PORT")); err != nil {
		panic(err)
	}
}
