package utils

import (
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/pkg/errors"
	"html"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gabrielopesantos/myDrive-api/internal/models"
)

type Claims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.StandardClaims
}

//func GenerateJWT(user *models.User, config *config.Config) (string, error) {
func GenerateJWT(user *models.User, config *config.Config) (string, error) {
	claims := &Claims{
		Email: user.Email,
		ID:    user.UserID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Server.JwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Extract JWT from request
func ExtractJWTFromRequest(r *http.Request) (map[string]interface{}, error) {
	// Get the JWT string
	tokenString := ExtractBearerToken(r)

	// Initialize a new instance of `Claims`
	claims := jwt.MapClaims{}

	// Parse JWT string and repos result in `claims`
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (jwtKey interface{}, err error) {
		return jwtKey, err
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// Extract bearer token from request Authorization header
func ExtractBearerToken(r *http.Request) string {
	headerAuthorization := r.Header.Get("Authorization")
	bearerToken := strings.Split(headerAuthorization, " ")
	return html.EscapeString(bearerToken[1])
}
