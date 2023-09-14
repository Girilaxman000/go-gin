package repository

import (
	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/database"
)

func OrdersCreate(order models.Orders) (err error) {
	var databaseOrder models.Orders
	databaseOrder.Name = order.Name
	return database.DB.Create(&databaseOrder).Error
}
