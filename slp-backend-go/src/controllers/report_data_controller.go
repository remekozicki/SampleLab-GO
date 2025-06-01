package controllers

import (
	"net/http"
	"samplelab-go/src/dto"
	"samplelab-go/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllReportData(c *gin.Context) {
	data, err := services.GetAllReportData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetReportDataBySampleID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("sampleId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawne ID próbki"})
		return
	}
	dtoData, err := services.GetReportDataBySampleID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dtoData)
}

func SaveReportData(c *gin.Context) {
	var input dto.ReportDataDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.SaveReportData(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func UpdateReportData(c *gin.Context) {
	SaveReportData(c)
}

func DeleteReportData(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawne ID"})
		return
	}
	if err := services.DeleteReportData(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
