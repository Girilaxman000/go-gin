package repository

import (
	"github.com/Girilaxman000/go-gin/api/dtos"
	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/commonfunctions"
	"github.com/Girilaxman000/go-gin/database"
	"golang.org/x/crypto/bcrypt"
)

func UsersCreate(user models.User) (err error) {
	var users models.User
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	users.Email = user.Email
	users.Password = string(hash)
	return database.DB.Create(&users).Error
}

func UsersLogin(user dtos.UserLoginDetail) (userToken string, err error) {
	//retrieve use from database
	var databaseUser models.User

	//steps needed for login
	// 1. first find if uername/email is present in database
	// 2. then match password
	// 3. authenticate user
	// 4. create jwt token with help of secret key
	// 5. send back to user

	err = database.DB.Find(&databaseUser, "email = ?", user.Email).Error

	if err != nil {
		return "", err
	}

	err = commonfunctions.HashPassword(databaseUser, user)
	if err != nil {
		return "", err // Return an empty string for userToken
	}

	tokenString, err := commonfunctions.GenerateToken(databaseUser)
	database.DB.Commit().Debug()
	return tokenString, err
}
