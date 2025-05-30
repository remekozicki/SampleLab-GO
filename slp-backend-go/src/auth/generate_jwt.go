package auth

import (
	"errors"
	"samplelab-go/src/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("twoj_sekretny_klucz")

func GenerateJWT(user models.DBUser) (string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role.String(), // zakładam, że masz Role jako enum z metodą String()
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

var ErrInvalidToken = errors.New("nieprawidłowy token")

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
