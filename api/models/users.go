package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"email" json:"email" validate:"required,email"`
	FullName string `gorm:"full_name" json:"full_name" validate:"required"`
	Phone    string `gorm:"phone" json:"phone"  validate:"required,phone"`
	Gender   string `gorm:"gender" json:"gender" validate:"required,gender"`
	Password string `gorm:"password" json:"password" validate:"required"`
}
