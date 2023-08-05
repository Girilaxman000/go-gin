package controllers

import (
	"net/http"

	"github.com/Girilaxman000/go-gin/api/services"
	"github.com/gin-gonic/gin"
)

func UsersCreate(c *gin.Context) {

}

func GetAllUsers(c *gin.Context) {
	Users := services.GetAllUsers()
	c.JSON(http.StatusOK, Users)
}
func GetOneUser(c *gin.Context) {

}

func UpdateOneUser(c *gin.Context) {

}

func DeleteOneUser(c *gin.Context) {

}
