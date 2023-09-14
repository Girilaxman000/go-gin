package services

import (
	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/api/repository"
)

func OrdersCreate(order models.Orders) (err error) {
	return repository.OrdersCreate(order)
}
