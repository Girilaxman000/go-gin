package controllers

import (
	"net/http"

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

func UsersLogin(c *gin.Context) {
	Users := services.UsersLogin()
	c.JSON(http.StatusOK, Users)
}
