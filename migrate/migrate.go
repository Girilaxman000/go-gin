package main

import (
	"github.com/Girilaxman000/go-gin/database"
	"github.com/Girilaxman000/go-gin/initializers"
	"github.com/Girilaxman000/go-gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	database.ConnectToDatabase()
}

func main() {
	//create table in database
	database.DB.AutoMigrate(&models.Post{})
}
