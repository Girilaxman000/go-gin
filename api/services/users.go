package services

import (
	"github.com/Girilaxman000/go-gin/api/dtos"
	"github.com/Girilaxman000/go-gin/api/models"
	"github.com/Girilaxman000/go-gin/api/repository"
)

func UsersCreate(user models.User) (err error) {
	return repository.UsersCreate(user)
}

func UsersLogin(user dtos.UserLoginDetail) (userToken string, err error) {
	return repository.UsersLogin(user)
}
