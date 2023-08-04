package controllers

import (
	"net/http"

	"github.com/Girilaxman000/go-gin/services"
	"github.com/gin-gonic/gin"
)

func ProductsCreate(c *gin.Context) {

}

func GetAllProducts(c *gin.Context) {
    products := services.GetAllProducts()
    c.JSON(http.StatusOK, products)
}
func GetOneProduct(c *gin.Context) {

}

func UpdateOneProduct(c *gin.Context) {

}

func DeleteOneProduct(c *gin.Context) {

}
