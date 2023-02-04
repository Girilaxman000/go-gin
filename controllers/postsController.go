package controllers

import (
	"github.com/Girilaxman000/go-gin/database"
	"github.com/Girilaxman000/go-gin/models"
	"github.com/gin-gonic/gin"
)

// make function name capital to use it on other files.
func PostsCreate(c *gin.Context) {
	//Get data off req body

	//Create a post
	post := models.Post{Title: "Hello", Body: "Post Body"}
	result := database.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}
