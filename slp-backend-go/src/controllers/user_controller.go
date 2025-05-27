package controllers

import (
	"errors"
	"net/http"
	"samplelab-go/src/models"
	"samplelab-go/src/services"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd przy pobieraniu użytkowników"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.RegisterUser(input)
	if err != nil {
		if err == services.ErrEmailTaken {
			c.JSON(http.StatusConflict, gin.H{"error": "Email już istnieje"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd serwera"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"role":  user.Role,
	})
}

func Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.AuthenticateUser(input.Email, input.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Nieprawidłowy e-mail lub hasło"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd serwera"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
