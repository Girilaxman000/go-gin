package controllers

import (
	"net/http"

	"github.com/Girilaxman000/go-gin/api/dtos"
	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/api/services"
	"github.com/gin-gonic/gin"
)

func UsersCreate(ctx *gin.Context) {
	var users models.User
	err := ctx.BindJSON(&users)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	err = services.UsersCreate(users)
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed",
			})
			return
		}
	}

}

func UsersLogin(ctx *gin.Context) {
	var user dtos.UserLoginDetail
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	userToken, err := services.UsersLogin(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": userToken,
	})

}
