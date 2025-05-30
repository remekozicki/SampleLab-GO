package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"samplelab-go/src/dto"
	"samplelab-go/src/services"
)

func GetClientList(c *gin.Context) {
	clients, err := services.GetAllClients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd przy pobieraniu klientów"})
		return
	}
	c.JSON(http.StatusOK, clients)
}

func AddClient(c *gin.Context) {
	var input dto.ClientDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.SaveClient(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie udało się zapisać klienta"})
		return
	}
	c.Status(http.StatusCreated)
}

func EditClient(c *gin.Context) {
	AddClient(c) // identyczna logika
}

func DeleteClient(c *gin.Context) {
	var req struct {
		ID uint `uri:"id" binding:"required"`
	}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawne ID"})
		return
	}

	if err := services.DeleteClient(req.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
