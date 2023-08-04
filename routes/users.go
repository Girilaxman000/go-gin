package routes

import (
	"github.com/Girilaxman000/go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine) {
	router.POST("/users", controllers.UsersCreate)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/users/:id", controllers.GetOneUser)
	router.PUT("/users/:id", controllers.UpdateOneUser)
	router.DELETE("/users/:id", controllers.DeleteOneUser)

}
