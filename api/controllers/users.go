package controllers

import (
	"fmt"
	"net/http"

	"github.com/Girilaxman000/go-gin/api/dtos"
	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/api/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UsersCreate(ctx *gin.Context) {
	var users models.User
	fmt.Println(users)
	err := ctx.BindJSON(&users)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	validate := validator.New()
	errVal := validate.Struct(users)
	if errVal != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errVal,
		})
		return
	}

	err = services.UsersCreate(users)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
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
