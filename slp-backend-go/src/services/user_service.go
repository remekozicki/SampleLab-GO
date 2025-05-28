package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"samplelab-go/src/db"
	"samplelab-go/src/dto"
	"samplelab-go/src/models"
)

func GetAllUsers() ([]models.DBUser, error) {
	conn := db.GetDB() // tu conn to *gorm.DB

	var users []models.DBUser
	result := conn.Find(&users)
	return users, result.Error
}

func FindUserByEmail(email string) (models.DBUser, error) {
	conn := db.GetDB() // <- to musi być *gorm.DB

	var user models.DBUser
	err := conn.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.DBUser{}, errors.New("użytkownik nie istnieje")
		}
		return models.DBUser{}, err
	}

	return user, nil
}

var ErrEmailTaken = errors.New("adres email jest już zajęty")

func RegisterUser(input dto.RegisterInput) (*models.DBUser, error) {
	conn := db.GetDB()

	// sprawdź, czy email już istnieje
	var existing models.DBUser
	if err := conn.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		return nil, ErrEmailTaken
	}

	password := generateRandomPassword(12)

	// haszuj hasło
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// zapisz użytkownika
	user := models.DBUser{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
		Role:     input.Role,
	}
	if err := conn.Create(&user).Error; err != nil {
		return nil, err
	}
	user.Password = password
	return &user, nil
}

func generateRandomPassword(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		// fallback gdyby coś poszło nie tak
		return "DefaultPass123"
	}
	return base64.StdEncoding.EncodeToString(bytes)[:length]
}

var ErrWrongPassword = errors.New("stare hasło jest nieprawidłowe")

func ChangePassword(email string, req dto.ChangePasswordRequest) error {
	conn := db.GetDB()

	var user models.DBUser
	if err := conn.Where("email = ?", email).First(&user).Error; err != nil {
		return errors.New("użytkownik nie istnieje")
	}

	// Sprawdź stare hasło
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return ErrWrongPassword
	}

	// Haszuj nowe hasło
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Zapisz nowe hasło
	user.Password = string(hashed)
	return conn.Save(&user).Error
}
