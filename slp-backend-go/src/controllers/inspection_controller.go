package controllers

import (
	"net/http"
	"samplelab-go/src/dto"
	"samplelab-go/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetInspectionList(c *gin.Context) {
	inspections, err := services.GetAllInspections()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie udało się pobrać listy"})
		return
	}
	c.JSON(http.StatusOK, inspections)
}

func AddInspection(c *gin.Context) {
	var input dto.InspectionDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawne dane"})
		return
	}
	if err := services.SaveInspection(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd zapisu"})
		return
	}
	c.Status(http.StatusOK)
}

func EditInspection(c *gin.Context) {
	var input dto.InspectionDto
	if err := c.ShouldBindJSON(&input); err != nil || input.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawne dane"})
		return
	}
	if err := services.UpdateInspection(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd aktualizacji"})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteInspection(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawne ID"})
		return
	}
	if err := services.DeleteInspection(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd usuwania"})
		return
	}
	c.Status(http.StatusOK)
}
