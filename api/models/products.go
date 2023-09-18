package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint     `gorm:"primaryKey"`
	Name        string   `gorm:"name" json:"name" validate:"required,name"`
	Description string   `gorm:"description" json:"description" validate:"required,description"`
	ImageURL    string   `gorm:"image_url" json:"image_url" validate:"required,image_url"`
	Price       float64  `gorm:"price" json:"price" validate:"required,price"`
	Orders      []Orders `gorm:"foreignKey:ProductId;references:ID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// TableName returns the custom table name for the Product model
// func (c *Product) TableName() string {
// 	return "products" // Change "products" to your desired table name
// }
