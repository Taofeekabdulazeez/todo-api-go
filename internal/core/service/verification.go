package service

import (
	"todo-api-go/internal/core/model"
	"todo-api-go/pkg/database"
)

type VerificationService struct {
}

func (s *VerificationService) InvalidateVerificationToken(token string) (model.Verification, error) {
	var verification model.Verification
	result := database.DB.First(&verification, "token = ? AND used = ?", token, false)
	verification.Used = true
	database.DB.Save(&verification)

	return verification, result.Error
}

func (s *VerificationService) CreateVerificationToken(email string) (string, error) {
	verification := &model.Verification{
		Token: "1234", // random string generation logic to be implemented
		Email: email,
	}

	result := database.DB.Create(verification)
	return verification.Token, result.Error
}
