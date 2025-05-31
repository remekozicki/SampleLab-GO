package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"samplelab-go/src/dto"
	"samplelab-go/src/services"
)

func GetAllCodes(c *gin.Context) {
	codes, err := services.GetAllCodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, codes)
}

func AddCode(c *gin.Context) {
	var input dto.CodeDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.SaveCode(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func EditCode(c *gin.Context) {
	AddCode(c) // identyczna logika
}

func DeleteCode(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Brak ID"})
		return
	}
	if err := services.DeleteCode(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
