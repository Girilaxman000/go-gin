package main

import (
	"github.com/Girilaxman000/go-gin/controllers"
	"github.com/Girilaxman000/go-gin/database"
	"github.com/Girilaxman000/go-gin/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDatabase()
}

func main() {
	router := gin.Default()
	router.POST("/posts", controllers.PostsCreate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsById)
	router.PUT("/posts/:id", controllers.PostsUpdate)
	router.Run()
}
