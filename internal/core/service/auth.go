package service

import (
	"errors"
	"strings"
	"todo-api-go/internal/core/model"
	"todo-api-go/internal/api/request"
	"todo-api-go/pkg/database"
	"todo-api-go/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (s *UserService) CreateUser(data request.CreateUserRequest) (*model.User, error) {

	data.Email = strings.TrimSpace(data.Email)
	data.Password = strings.TrimSpace(data.Password)

	user := model.User{Email: data.Email, Password: data.Password}

	hash, err := utils.HashPassword(user.Password)

	if err != nil {
		return nil, errors.New("error hashing user password")
	}

	user.Password = hash

	result := database.DB.Create(&user)

	if result.Error != nil {
		return nil, errors.New("error creating user")
	}

	return &user, nil
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	result := database.DB.First(&user, "email = ?", email)

	return &user, result.Error
}

func (s *UserService) VerifyUser(email string, pasword string) (bool, *model.User, error) {
	var user model.User

	result := database.DB.First(&user, "email = ?", email)

	if result.Error != nil {
		return false, nil, errors.New("error getting user")
	}

	success := s.VerifyPassword(pasword, user.Password)

	return success, &user, nil
}

func (s *UserService) VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
