package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint     `gorm:"primaryKey"`
	Name        string   `gorm:"column:name" json:"name" validate:"required"`
	Description string   `gorm:"column:description" json:"description" validate:"required"`
	ImageURL    string   `gorm:"column:image_url" json:"image_url" validate:"required,url"`
	Price       float64  `gorm:"column:price" json:"price" validate:"required,gt=0"` //  price should be greater than 0
	Orders      []Orders `gorm:"foreignKey:ProductId;references:ID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func ValidateProduct(product Product) error {
	return validate.Struct(product)
}
