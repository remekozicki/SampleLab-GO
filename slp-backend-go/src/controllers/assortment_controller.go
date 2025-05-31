package controllers

import (
	"net/http"
	"samplelab-go/src/dto"
	"samplelab-go/src/services"

	"github.com/gin-gonic/gin"
)

func GetAssortmentList(c *gin.Context) {
	list := services.GetAllAssortments()
	c.JSON(http.StatusOK, list)
}

func AddAssortment(c *gin.Context) {
	var input dto.AssortmentDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.SaveAssortment(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd zapisu asortymentu"})
		return
	}
	c.Status(http.StatusOK)
}

func EditAssortment(c *gin.Context) {
	var input dto.AssortmentDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateAssortment(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd aktualizacji asortymentu"})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteAssortment(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteAssortmentByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd usuwania"})
		return
	}
	c.Status(http.StatusOK)
}
