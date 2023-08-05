package models

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"name" json:"name" validate:"required,name"`
	Description string `gorm:"description" json:"description" validate:"required,description"`
}

// TableName returns the custom table name for the Product model
// func (c *Product) TableName() string {
// 	return "products" // Change "products" to your desired table name
// }

// Price       float64 `gorm:"price" json:"price" validate:"required,price"`
// 	Stock       int     `gorm:"stock" json:"stock" validate:"required,stock"`
// 	Category    string  `gorm:"category" json:"category" validate:"required,category"`
// 	Brand       string  `gorm:"brand" json:"brand" `
// 	// ImageURL    string  `gorm:"image_url" json:"image_url" validate:"required,image_url"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
