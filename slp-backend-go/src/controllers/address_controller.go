package controllers

import (
	"net/http"
	"samplelab-go/src/services"

	"github.com/gin-gonic/gin"
)

func GetAddressList(c *gin.Context) {
	addresses, err := services.GetAllAddresses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Błąd przy pobieraniu adresów"})
		return
	}
	c.JSON(http.StatusOK, addresses)
}
