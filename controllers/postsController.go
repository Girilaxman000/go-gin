package controllers

import (
	"github.com/Girilaxman000/go-gin/database"
	"github.com/Girilaxman000/go-gin/models"
	"github.com/gin-gonic/gin"
)

// make function name capital to use it on other files.
func PostsCreate(c *gin.Context) {
	//Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
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

func PostsIndex(c *gin.Context) {

	posts := []models.Post{}
	//finding all posts and assigning to above variable
	database.DB.Find(&posts)
	//respond with
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsById(c *gin.Context) {
	//Get id of url
	//getting an id
	id := c.Param("id")
	post := models.Post{}

	//database.DB.First(models.Post{}, id)

	//first argument passing referecnce to post variable
	database.DB.First(&post, id)
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//Get id of url
	id := c.Param("id")

	//Get the data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Find the post we were updating
	post := models.Post{}
	database.DB.First(&post, id)

	//update it
	database.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	//respond with id
	c.JSON(200, gin.H{
		"post": post,
	})
}
