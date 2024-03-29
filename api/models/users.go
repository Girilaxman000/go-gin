package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	ID        uint     `gorm:"primaryKey"`
	Email     string   `gorm:"column:email" json:"email" validate:"required,email"`
	Password  string   `gorm:"column:password" json:"password" validate:"required"`
	Orders    []Orders `gorm:"foreignKey:UserId;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// IsAdmin  bool   `gorm:"default:false" json:"isAdmin"`

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateUser(user User) error {
	return validate.Struct(user)
}
