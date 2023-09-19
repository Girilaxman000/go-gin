package controllers

import (
	"net/http"
	"strconv"

	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/api/services"
	"github.com/gin-gonic/gin"
)

func OrdersCreate(ctx *gin.Context) {
	var order models.Orders
	//1st bindjson
	errBind := ctx.BindJSON(&order)
	if errBind != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errBind.Error(),
		})
		return
	}
	//2nd bind custom data
	userID := ctx.MustGet("user")
	userIDStr, ok := userID.(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID type",
		})
		return
	}

	uintVal, errs := strconv.ParseUint(userIDStr, 10, 64)
	if errs != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID format",
		})
		return
	}
	order.UserId = uint(uintVal)

	err := services.OrdersCreate(order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}
