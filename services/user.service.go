package services

import (
	"errors"
	"strings"
	"todo-api-go/database"
	"todo-api-go/models"
	"todo-api-go/requests"
	"todo-api-go/utils"
)

type UserService struct{}

func (s *UserService) CreateUser(data requests.CreateUserRequest) (*models.User, error) {

	data.Email = strings.TrimSpace(data.Email)
	data.Password = strings.TrimSpace(data.Password)

	user := models.User{Email: data.Email, Password: data.Password}

	hash, err := utils.HashPassword(user.Password)

	if err != nil {
		return nil, errors.New("error hashing password")
	}

	user.Password = hash

	result := database.DB.Create(&user)

	if result.Error != nil {
		return nil, errors.New("error creating user")
	}

	return &user, nil
}
