package infrastructure

import (
	"task_management_api_with_clean_architecture/usecase"

	"github.com/dgrijalva/jwt-go"
)

type TokenService struct {
	SecretKey string
}

func NewTokenService(secretKey string) usecase.ITokenService {
	return &TokenService{SecretKey: secretKey}
}

func (ts *TokenService) GenerateToken(id string, username string, role string, expiryDate int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"role":     role,
		"exp":      expiryDate,
	})

	jwtToken, err := token.SignedString([]byte(ts.SecretKey))

	return jwtToken, err
}
