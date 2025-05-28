package controllers

import (
	"errors"
	"net/http"
	"samplelab-go/src/dto"
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
	var input dto.RegisterInput
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
		"id":       user.ID,
		"email":    user.Email,
		"name":     user.Name,
		"role":     user.Role,
		"password": user.Password,
	})
}

func Login(c *gin.Context) {
	var input dto.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := services.AuthenticateUser(input.Email, input.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Nieprawidłowy e-mail lub hasło"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd serwera"})
		}
		return
	}

	c.JSON(http.StatusOK, resp)
}

func ChangePassword(c *gin.Context) {
	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawne dane"})
		return
	}

	email := c.MustGet("email").(string)
	if email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Brak emaila w nagłówku"})
		return
	}

	err := services.ChangePassword(email, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hasło zostało zmienione"})
}

func ChangePasswordByAdmin(c *gin.Context) {
	var req dto.ChangePasswordRequest
	email := c.Param("email")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawne dane"})
		return
	}

	err := services.ChangePasswordByAdmin(email, req.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hasło zostało zmienione przez administratora"})
}

func DeleteUserByEmail(c *gin.Context) {
	email := c.Param("email")

	err := services.DeleteUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Użytkownik został usunięty"})
}
