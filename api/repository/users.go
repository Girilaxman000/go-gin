package repository

import (
	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/database"
	"golang.org/x/crypto/bcrypt"
)

func UsersCreate(user models.User) (err error) {
	var users models.User

	hash, errorC := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errorC != nil {
		return errorC
	}
	users.Email = user.Email
	users.FullName = user.FullName
	users.Gender = user.Gender
	users.Password = string(hash)
	users.Phone = user.Phone
	return database.DB.Create(&user).Error
}

func UsersLogin() string {
	return "This returns users"
}
