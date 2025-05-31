package controllers

import (
	"net/http"
	"samplelab-go/src/dto"
	"samplelab-go/src/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProductGroupList(c *gin.Context) {
	groups, err := services.GetAllProductGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Nie można pobrać grup"})
		return
	}
	c.JSON(http.StatusOK, groups)
}

func AddProductGroup(c *gin.Context) {
	var input dto.ProductGroupSaveDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawne dane"})
		return
	}
	if err := services.SaveProductGroup(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd zapisu"})
		return
	}
	c.Status(http.StatusOK)
}

func EditProductGroup(c *gin.Context) {
	var input dto.ProductGroupDto
	if err := c.ShouldBindJSON(&input); err != nil || input.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawne dane"})
		return
	}
	if err := services.UpdateProductGroup(input.ID, dto.ProductGroupSaveDto{Name: input.Name}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd aktualizacji"})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteProductGroup(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Niepoprawne ID"})
		return
	}
	if err := services.DeleteProductGroup(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd usuwania"})
		return
	}
	c.Status(http.StatusOK)
}
