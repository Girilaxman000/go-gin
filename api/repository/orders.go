package repository

import (
	"fmt"

	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/database"
)

func OrdersCreate(order models.Orders) error {
	fmt.Println("__create database")
	return database.DB.Create(&order).Error
}
