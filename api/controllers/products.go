package controllers

import (
	"net/http"

	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/api/services"
	"github.com/gin-gonic/gin"
)

func ProductsCreate(ctx *gin.Context) {
	//define data types
	//here can be defined dtos this is from users request

	var product models.Product
	// Bind user input to the struct
	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// validate product
	errVal := models.ValidateProduct(product)
	if errVal != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errVal.Error(),
		})
		return
	}

	//pass to services
	err := services.ProductsCreate(product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not create products",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"products": "Products created successfully",
	})
}

func GetAllProducts(ctx *gin.Context) {
	productsList, err := services.GetAllProducts()
	if err != nil {

	}
	ctx.JSON(http.StatusOK, productsList)
}
func GetOneProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := services.GetOneProduct(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func UpdateOneProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var products models.Product
	if err := ctx.BindJSON(&products); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	errVal := models.ValidateProduct(products)
	if errVal != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errVal.Error(),
		})
		return
	}
	err := services.UpdateOneProduct(id, products)
	if err != nil {

	}

}

func DeleteOneProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	err := services.DeleteOneProduct(id)
	if err != nil {

	}

}
