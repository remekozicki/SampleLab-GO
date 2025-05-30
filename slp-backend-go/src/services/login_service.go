package services

import (
	"errors"
	"samplelab-go/src/auth"
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
	"samplelab-go/src/models"

	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("nieprawidłowy e-mail lub hasło")

func AuthenticateUser(email, password string) (*dto.LoginResponse, error) {
	conn := db.GetDB()
	var user models.User

	if err := conn.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := auth.GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	resp := &dto.LoginResponse{
		User: dto.User{
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
		Token: token,
	}

	return resp, nil
}
