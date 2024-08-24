package main

import (
	"go-api/config"
	"go-api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	r := gin.Default()

	routes.RegisterRoutes(r)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
