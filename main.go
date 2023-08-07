package main

import (
	"github.com/Girilaxman000/go-gin/api/routes"
	"github.com/Girilaxman000/go-gin/database"
	"github.com/Girilaxman000/go-gin/initializers"
	"github.com/Girilaxman000/go-gin/migrate"
	"github.com/gin-gonic/gin"
)

// this function runs right before main
func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDatabase()
	migrate.SyncDatabase()

}

func main() {
	router := gin.Default()
	routes.ProductsRoutes(router)
	routes.UsersRoutes(router)
	router.Run(":3000")
}
