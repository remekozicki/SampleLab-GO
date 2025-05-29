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
	users.POST("/login", controllers.Login)
	users.Use(auth.JWTMiddleware())
	{
		users.GET("/", auth.RequireAdminRole(), controllers.GetAllUsers)
		users.POST("/register", controllers.Register)
		users.POST("/change-password", controllers.ChangePassword)
		users.POST("/change-password/:email", auth.RequireAdminRole(), controllers.ChangePasswordByAdmin)
		users.DELETE("/:email", auth.RequireAdminRole(), controllers.DeleteUserByEmail)

	}

	address := r.Group("/address")
	address.Use(auth.JWTMiddleware())
	{
		address.GET("/list", controllers.GetAddressList)
	}

	// Routing

	err := r.Run(":8090")
	if err != nil {
		return
	} // Nasłuchuj na porcie 8080
}
