package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"samplelab-go/src/auth"
	"samplelab-go/src/controllers"
	"samplelab-go/src/db"
)

func main() {
	// 1. Połączenie z bazą
	db.InitDB()

	// 2. Router
	r := gin.Default()

	// Endpoint testowy
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "backend running!"})
	})

	db.InitDB()

	users := r.Group("/users")
	users.Use(auth.JWTMiddleware())
	{
		users.GET("/", controllers.GetAllUsers)
		users.POST("/register", controllers.Register)
		users.POST("/change-password", controllers.ChangePassword)
		// dodaj pozostałe: deleteUser, changePasswordByAdmin itp.
	}

	// Routing

	r.POST("/users/login", controllers.Login)
	err := r.Run(":8080")
	if err != nil {
		return
	} // Nasłuchuj na porcie 8080
}
