package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint     `gorm:"primaryKey"`
	Email     string   `gorm:"email" json:"email" validate:"required,email"`
	Password  string   `gorm:"password" json:"password" validate:"required"`
	Orders    []Orders `gorm:"foreignKey:UserId;references:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// IsAdmin  bool   `gorm:"default:false" json:"isAdmin"`
