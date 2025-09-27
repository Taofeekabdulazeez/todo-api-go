package service

import (
	"time"
	"todo-api-go/pkg/config"

	"github.com/golang-jwt/jwt/v5"
)

type VerificationService struct{}

func (s *VerificationService) CreateToken(value any, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"value": value,
			"exp":   time.Now().Add(expiresIn).Unix(),
		})

	tokenString, err := token.SignedString(config.JWT_SECRET)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *VerificationService) VerifyToken(tokenString string) (any, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return config.JWT_SECRET, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", jwt.ErrTokenInvalidId
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", jwt.ErrTokenInvalidClaims
	}

	value, ok := claims["value"]
	if !ok {
		return "", jwt.ErrInvalidKey
	}

	return value, nil
}
