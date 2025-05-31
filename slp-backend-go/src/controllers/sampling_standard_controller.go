package controllers

import (
	"net/http"
	"samplelab-go/src/dto"
	"samplelab-go/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSamplingStandardList(c *gin.Context) {
	standards, err := services.GetAllSamplingStandards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie można pobrać listy"})
		return
	}
	c.JSON(http.StatusOK, standards)
}

func AddSamplingStandard(c *gin.Context) {
	var input dto.SamplingStandardDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawne dane"})
		return
	}
	if err := services.SaveSamplingStandard(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd zapisu"})
		return
	}
	c.Status(http.StatusOK)
}

func EditSamplingStandard(c *gin.Context) {
	var input dto.SamplingStandardDto
	if err := c.ShouldBindJSON(&input); err != nil || input.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawne dane"})
		return
	}
	if err := services.UpdateSamplingStandard(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd aktualizacji"})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteSamplingStandard(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawne ID"})
		return
	}
	if err := services.DeleteSamplingStandard(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd usuwania"})
		return
	}
	c.Status(http.StatusOK)
}
