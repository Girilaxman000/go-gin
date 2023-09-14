package routes

import (
	"github.com/Girilaxman000/go-gin/api/controllers"
	"github.com/Girilaxman000/go-gin/middleware"
	"github.com/gin-gonic/gin"
)

func OrdersRoutes(router *gin.Engine) {
	router.POST("/orders", middleware.AuthMiddleware, controllers.OrdersCreate)
}
