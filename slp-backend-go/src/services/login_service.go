package services

import (
	"errors"
	"gorm.io/gorm"
	"samplelab-go/src/auth"
	"samplelab-go/src/db"
	"samplelab-go/src/models"

	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("nieprawidłowy e-mail lub hasło")

func AuthenticateUser(email, password string) (string, error) {
	conn := db.GetDB() // <- musi być *gorm.DB

	var user models.DBUser
	err := conn.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrInvalidCredentials
		}
		return "", err
	}

	// Sprawdzenie hasła
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	// JWT
	token, err := auth.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
