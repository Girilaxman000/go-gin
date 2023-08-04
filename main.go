package main

import (
	"github.com/Girilaxman000/go-gin/database"
	"github.com/Girilaxman000/go-gin/initializers"
	"github.com/Girilaxman000/go-gin/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDatabase()
}

func main() {
	router := gin.Default()
	routes.ProductsRoutes(router)
	routes.UsersRoutes(router)
	router.Run(":3000")
}
