package models

type Orders struct {
	ID        uint `gorm:"primaryKey"`
	ProductId uint
	UserId    uint
	Cost      int
	Quantity  int
}
