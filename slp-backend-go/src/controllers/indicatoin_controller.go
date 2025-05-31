package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"samplelab-go/src/dto"
	"samplelab-go/src/services"
)

func GetAllIndications(c *gin.Context) {
	list, err := services.GetAllIndications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd pobierania danych"})
		return
	}
	c.JSON(http.StatusOK, list)
}

func GetIndicationByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowe ID"})
		return
	}

	indication, err := services.GetIndicationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nie znaleziono wskazania"})
		return
	}

	c.JSON(http.StatusOK, indication)
}

func SaveIndication(c *gin.Context) {
	var input dto.IndicationDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.SaveIndication(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie udało się zapisać"})
		return
	}
	c.Status(http.StatusCreated)
}

func EditIndication(c *gin.Context) {
	SaveIndication(c)
}

func DeleteIndication(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowe ID"})
		return
	}

	if err := services.DeleteIndication(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd usuwania wskazania"})
		return
	}
	c.Status(http.StatusOK)
}

//func GetIndicationsForSample(c *gin.Context) {
//	sampleIDStr := c.Param("sampleId")
//	sampleID, err := strconv.ParseInt(sampleIDStr, 10, 64)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowe ID próbki"})
//		return
//	}
//
//	indications, err := services.SelectIndicationsForSample(sampleID)
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, indications)
//}
