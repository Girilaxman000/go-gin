package main

import (
	"github.com/Girilaxman000/go-gin/api/routes"
	"github.com/Girilaxman000/go-gin/database"
	"github.com/Girilaxman000/go-gin/initializers"
	"github.com/Girilaxman000/go-gin/migrate"
	"github.com/gin-contrib/cors"
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

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	routes.ProductsRoutes(router)
	routes.UsersRoutes(router)
	routes.OrdersRoutes(router)
	router.Run(":8000")
}
