package repository

import (
	"os"
	"time"

	"github.com/Girilaxman000/go-gin/api/dtos"
	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/database"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func UsersCreate(user models.User) (err error) {
	var users models.User
	hash, errorC := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errorC != nil {
		return errorC
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

	// step1
	err = database.DB.Find(&databaseUser, "email = ?", user.Email).Error

	if err != nil {
		return "", err // Return an empty string for userToken
	}
	// step2
	err = bcrypt.CompareHashAndPassword([]byte(databaseUser.Password), []byte(user.Password))
	if err != nil {
		return "", err // Return an empty string for userToken
	}

	//for custom data to be  stored in jwt
	claims := jwt.MapClaims{
		"sub":      databaseUser.ID,
		"fullName": databaseUser.Email,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, errList := token.SignedString([]byte(os.Getenv("SECRET")))

	return tokenString, errList
}
