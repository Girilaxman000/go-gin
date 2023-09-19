package controllers

import (
	"net/http"

	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/api/services"
	"github.com/gin-gonic/gin"
)

func OrdersCreate(ctx *gin.Context) {
	// fmt.Println(ctx.MustGet("user"))

	var order models.Orders

	errBind := ctx.BindJSON(&order)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errBind.Error(),
		})
		return
	}
	err := services.OrdersCreate(order)
	if err != nil {
		return
	}
}
