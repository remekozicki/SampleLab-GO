package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"samplelab-go/src/models"
	"time"
)

var jwtKey = []byte("twoj_sekretny_klucz")

func GenerateJWT(user models.DBUser) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   user.Email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

var ErrInvalidToken = errors.New("nieprawidłowy token")

func ValidateJWT(tokenString string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
