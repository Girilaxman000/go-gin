package controllers

import (
	"fmt"

	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/api/services"
	"github.com/gin-gonic/gin"
)

func OrdersCreate(ctx *gin.Context) {
	fmt.Println(ctx.MustGet("user"))

	var order models.Orders

	err := services.OrdersCreate(order)
	if err != nil {
		return
	}
}
