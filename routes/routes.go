package routes

import (
	"go-api/controllers"
	"go-api/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the application's routes.
func RegisterRoutes(r *gin.Engine) {

	r.Use(middleware.AuthMiddleware())

	r.POST("/bid", controllers.BidHandler)
}
