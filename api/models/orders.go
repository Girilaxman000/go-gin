package models

type Orders struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"name" json:"name" validate:"required,name"`
	ProductId uint
	UserId    uint
}
