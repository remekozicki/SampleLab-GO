package controllers

import (
	"net/http"
	"samplelab-go/src/enum"
	"strconv"

	"github.com/gin-gonic/gin"
	"samplelab-go/src/dto"
	"samplelab-go/src/services"
)

func GetAllSamples(c *gin.Context) {
	samples, err := services.GetAllSamples()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, samples)
}

func GetSampleByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	sample, err := services.GetSampleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sample not found"})
		return
	}
	c.JSON(http.StatusOK, sample)
}

func SaveSample(c *gin.Context) {
	var sampleDto dto.SampleDto
	if err := c.ShouldBindJSON(&sampleDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.SaveSample(sampleDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func UpdateSample(c *gin.Context) {
	var sampleDto dto.SampleDto
	if err := c.ShouldBindJSON(&sampleDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.UpdateSample(sampleDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteSample(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = services.DeleteSample(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func FilterSamplesHandler(c *gin.Context) {
	var filter dto.SampleFilterDto
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	samples, total, err := services.FilterSamples(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd filtrowania próbek"})
		return
	}

	totalPages := int((total + int64(filter.PageSize) - 1) / int64(filter.PageSize))

	c.JSON(http.StatusOK, gin.H{
		"totalPages": totalPages,
		"samples":    samples,
	})
}

func GetFilters(c *gin.Context) {
	filters, err := services.GetFilters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd pobierania filtrów"})
		return
	}
	c.JSON(http.StatusOK, filters)
}

func CountSamples(c *gin.Context) {
	count, err := services.CountSamples()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie można policzyć próbek"})
		return
	}
	c.JSON(http.StatusOK, count)
}

func UpdateSampleStatus(c *gin.Context) {
	sampleID, err := strconv.ParseInt(c.Param("sampleId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sample ID"})
		return
	}

	statusParam := c.Param("status")
	progress, err := enum.ConvertProgressStatus(statusParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := services.UpdateSampleStatus(uint(sampleID), progress)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}
