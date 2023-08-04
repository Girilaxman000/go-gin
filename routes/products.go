package routes

import (
	"github.com/Girilaxman000/go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func ProductsRoutes(router *gin.Engine) {
	router.POST("/products", controllers.ProductsCreate)
	router.GET("/products", controllers.GetAllProducts)
	router.GET("/products/:id", controllers.GetOneProduct)
	router.PUT("/products/:id", controllers.UpdateOneProduct)
	router.DELETE("/products/:id", controllers.DeleteOneProduct)

}
