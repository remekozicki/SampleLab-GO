package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"samplelab-go/src/db"
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

func RegisterUser(input models.User) (*models.DBUser, error) {
	conn := db.GetDB()

	// sprawdź, czy email już istnieje
	var existing models.DBUser
	if err := conn.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		return nil, ErrEmailTaken
	}

	// haszuj hasło
	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
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

	return &user, nil
}
