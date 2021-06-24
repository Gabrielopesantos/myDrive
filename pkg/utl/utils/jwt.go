package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gabrielopesantos/myDrive-api/pkg/utl/models"
)

type Claims struct{
	Email string `json:"email"`
	ID string `json:"id"`
	jwt.StandardClaims
}


//func GenerateJWT(user *models.User, config *config.Config) (string, error) {
func GenerateJWT(user *models.User) (string, error) {
	claims := &Claims{
		Email: user.Email,
		ID: user.UserID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}