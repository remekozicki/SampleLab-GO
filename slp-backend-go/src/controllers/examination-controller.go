package controllers

import (
	"net/http"
	"samplelab-go/src/auth"
	"samplelab-go/src/dto"

	//"samplelab-go/src/dto"
	"samplelab-go/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetExaminationsBySampleID(c *gin.Context) {
	sampleID, err := strconv.ParseInt(c.Param("sampleId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowe ID próbki"})
		return
	}
	result, _ := services.GetAllExaminationsForSample(sampleID)
	c.JSON(http.StatusOK, result)
}

func GetExaminationByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowe ID"})
		return
	}
	result, err := services.GetExaminationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func SaveExamination(c *gin.Context) {
	if !auth.HasMinRole(c, "WORKER") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Brak uprawnień"})
		return
	}

	var dto dto.ExaminationDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.SaveExamination(dto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie udało się zapisać badania"})
		return
	}
	c.Status(http.StatusCreated)
}

func UpdateExamination(c *gin.Context) {
	if !auth.HasMinRole(c, "WORKER") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Brak uprawnień"})
		return
	}

	var dto dto.ExaminationDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.SaveExamination(dto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie udało się zaktualizować badania"})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteExamination(c *gin.Context) {
	if !auth.HasMinRole(c, "WORKER") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Brak uprawnień"})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nieprawidłowe ID"})
		return
	}

	if err := services.DeleteExamination(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
