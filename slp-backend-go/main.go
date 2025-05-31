package main

import (
	"net/http"
	"samplelab-go/src/auth"
	"samplelab-go/src/controllers"
	"samplelab-go/src/db"

	"github.com/gin-gonic/gin"
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
		users.GET("/", auth.RequireMinRole("ADMIN"), controllers.GetAllUsers)
		users.POST("/register", controllers.Register)
		users.POST("/change-password", controllers.ChangePassword)
		users.POST("/change-password/:email", auth.RequireMinRole("ADMIN"), controllers.ChangePasswordByAdmin)
		users.DELETE("/:email", auth.RequireMinRole("ADMIN"), controllers.DeleteUserByEmail)

	}

	address := r.Group("/address")
	address.Use(auth.JWTMiddleware())
	{
		address.GET("/list", controllers.GetAddressList)
	}

	clients := r.Group("/client")
	clients.Use(auth.JWTMiddleware())
	{
		clients.GET("/list", controllers.GetClientList)
		clients.POST("/save", auth.RequireMinRole("WORKER"), controllers.AddClient)
		clients.PUT("/update", auth.RequireMinRole("WORKER"), controllers.EditClient)
		clients.DELETE("/delete/:id", auth.RequireMinRole("WORKER"), controllers.DeleteClient)
	}

	assortment := r.Group("/assortment")
	assortment.Use(auth.JWTMiddleware())
	{
		assortment.GET("/list", controllers.GetAssortmentList)
		assortment.POST("/save", auth.RequireMinRole("WORKER"), controllers.AddAssortment)
		assortment.PUT("/update", auth.RequireMinRole("WORKER"), controllers.EditAssortment)
		assortment.DELETE("/delete/:id", auth.RequireMinRole("WORKER"), controllers.DeleteAssortment)
	}
	err := r.Run(":8090")
	if err != nil {
		return
	}
}
