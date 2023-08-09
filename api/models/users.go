package models

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"email" json:"email" validate:"required,email"`

	Password string `gorm:"password" json:"password" validate:"required"`
}
