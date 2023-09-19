package routes

import (
	"github.com/Girilaxman000/go-gin/api/controllers"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.UsersCreate)
	router.POST("/signin", controllers.UsersLogin)
}
