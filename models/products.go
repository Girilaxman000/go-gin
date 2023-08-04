package models

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Price       float64
	Stock       int
	Category    string
	Brand       string
	ImageURL    string
}
