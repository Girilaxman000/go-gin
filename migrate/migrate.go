package migrate

import (
	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/database"
)

func SyncDatabase() {
	//take dataabase reference and generate table using structs defined in models
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Product{})
}
