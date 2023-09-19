package models

import "fmt"

type Orders struct {
	ID        uint `gorm:"primaryKey"`
	ProductId uint
	UserId    uint
}

func ValidateOrders(order Orders) error {
	fmt.Println("valiadte")
	return validate.Struct(order)
}
