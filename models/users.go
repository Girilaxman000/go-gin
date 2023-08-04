package models

type User struct {
	ID           uint   `gorm:"primaryKey"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	ShippingAddr string
	BillingAddr  string
	ContactNum   string
}
