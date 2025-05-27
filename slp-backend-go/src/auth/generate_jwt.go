package auth

import (
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
