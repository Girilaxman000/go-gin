package repository

import (
	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/database"
)

func ProductsCreate(products models.Product) (err error) {
	// Retrieve the existing product from the database
	var product models.Product
	//Update value inside database table
	product.Name = products.Name
	product.Description = products.Description
	//finally save it
	return database.DB.Create(&product).Error
}
func GetAllProducts() (product []models.Product, err error) {
	var products []models.Product
	err = database.DB.Find(&products).Error
	return products, err
}

func GetOneProduct(id string) (product models.Product, err error) {
	var oneProduct models.Product
	err = database.DB.Where("id = ?", id).First(&oneProduct).Error
	return oneProduct, err
}

func UpdateOneProduct(id string, body models.Product) (err error) {
	// Retrieve the existing product from the database
	var existingProduct models.Product
	err = database.DB.Where("id = ?", id).First(&existingProduct).Error
	if err != nil {
		return err
	}

	// Update the fields of the existing product with the new values
	existingProduct.Name = body.Name
	existingProduct.Description = body.Description
	// Update other fields as needed

	// Save the updated product back to the database
	err = database.DB.Save(&existingProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteOneProduct(id string) (err error) {
	var existingProduct models.Product
	return database.DB.Where("id = ?", id).Unscoped().Delete(&existingProduct).Error
}
