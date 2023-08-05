package services

import (
	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/api/repository"
)

func ProductsCreate(products models.Product) (err error) {
	return repository.ProductsCreate(products)
}

func GetAllProducts() (product []models.Product, err error) {
	return repository.GetAllProducts()
}

func GetOneProduct(id string) (product models.Product, err error) {
	return repository.GetOneProduct(id)
}

func UpdateOneProduct(id string, body models.Product) (err error) {
	return repository.UpdateOneProduct(id, body)
}

func DeleteOneProduct(id string) (err error) {
	return repository.DeleteOneProduct(id)
}
