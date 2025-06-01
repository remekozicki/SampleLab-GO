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
	{
		users.POST("/login", controllers.Login)
	}

	users.Use(auth.JWTMiddleware())
	{
		users.GET("", auth.RequireMinRole("ADMIN"), controllers.GetAllUsers)
		users.POST("", controllers.Register)
		users.POST("/change-password", controllers.ChangePassword)
		users.POST("/change-password/:email", auth.RequireMinRole("ADMIN"), controllers.ChangePasswordByAdmin)
		users.DELETE("/:email", auth.RequireMinRole("ADMIN"), controllers.DeleteUserByEmail)
	}

	addresses := r.Group("/addresses")
	addresses.Use(auth.JWTMiddleware())
	{
		addresses.GET("", controllers.GetAddressList)
	}
	clients := r.Group("/client")
	clients.Use(auth.JWTMiddleware())
	{
		clients.GET("", controllers.GetClientList)
		clients.POST("", auth.RequireMinRole("WORKER"), controllers.AddClient)
		clients.PUT("/:id", auth.RequireMinRole("WORKER"), controllers.EditClient)
		clients.DELETE("/:id", auth.RequireMinRole("WORKER"), controllers.DeleteClient)
	}

	indications := r.Group("/indications")
	indications.Use(auth.JWTMiddleware())
	{
		indications.GET("", controllers.GetAllIndications)
		indications.GET("/:id", controllers.GetIndicationByID)
		indications.POST("", auth.RequireMinRole("WORKER"), controllers.SaveIndication)
		indications.PUT("/:id", auth.RequireMinRole("WORKER"), controllers.EditIndication)
		indications.DELETE("/:id", auth.RequireMinRole("WORKER"), controllers.DeleteIndication)
	}

	assortments := r.Group("/assortments")
	assortments.Use(auth.JWTMiddleware())
	{
		assortments.GET("", controllers.GetAssortmentList)
		assortments.POST("", auth.RequireMinRole("WORKER"), controllers.AddAssortment)
		assortments.PUT("/:id", auth.RequireMinRole("WORKER"), controllers.EditAssortment)
		assortments.DELETE("/:id", auth.RequireMinRole("WORKER"), controllers.DeleteAssortment)
	}

	productGroups := r.Group("/product-groups")
	productGroups.Use(auth.JWTMiddleware())
	{
		productGroups.GET("", controllers.GetProductGroupList)
		productGroups.POST("", auth.RequireMinRole("WORKER"), controllers.AddProductGroup)
		productGroups.PUT("/:id", auth.RequireMinRole("WORKER"), controllers.EditProductGroup)
		productGroups.DELETE("/:id", auth.RequireMinRole("WORKER"), controllers.DeleteProductGroup)
	}

	samplingStandards := r.Group("/sampling-standards")
	samplingStandards.Use(auth.JWTMiddleware())
	{
		samplingStandards.GET("", controllers.GetSamplingStandardList)
		samplingStandards.POST("", auth.RequireMinRole("WORKER"), controllers.AddSamplingStandard)
		samplingStandards.PUT("/:id", auth.RequireMinRole("WORKER"), controllers.EditSamplingStandard)
		samplingStandards.DELETE("/:id", auth.RequireMinRole("WORKER"), controllers.DeleteSamplingStandard)
	}

	codes := r.Group("/codes")
	codes.Use(auth.JWTMiddleware())
	{
		codes.GET("", controllers.GetAllCodes)
		codes.POST("", auth.RequireMinRole("WORKER"), controllers.AddCode)
		codes.PUT("/:id", auth.RequireMinRole("WORKER"), controllers.EditCode)
		codes.DELETE("/:id", auth.RequireMinRole("WORKER"), controllers.DeleteCode)
	}

	inspections := r.Group("/inspections")
	inspections.Use(auth.JWTMiddleware())
	{
		inspections.GET("", controllers.GetInspectionList)
		inspections.POST("", auth.RequireMinRole("WORKER"), controllers.AddInspection)
		inspections.PUT("/:id", auth.RequireMinRole("WORKER"), controllers.EditInspection)
		inspections.DELETE("/:id", auth.RequireMinRole("WORKER"), controllers.DeleteInspection)
	}

	reportData := r.Group("/report-data")
	reportData.Use(auth.JWTMiddleware())

	{
		reportData.GET("", controllers.GetAllReportData)
		reportData.GET("/:sampleId", controllers.GetReportDataBySampleID)
		reportData.POST("", auth.RequireMinRole("WORKER"), controllers.SaveReportData)
		reportData.PUT("", auth.RequireMinRole("WORKER"), controllers.UpdateReportData)
		reportData.DELETE("/:id", auth.RequireMinRole("WORKER"), controllers.DeleteReportData)
	}

	examination := r.Group("/examination")
	examination.Use(auth.JWTMiddleware())

	{
		//examination.GET("/sample/:sampleId", controllers.GetExaminationsBySampleID)
		examination.GET("/:id", controllers.GetExaminationByID)
		//examination.POST("", auth.RequireMinRole("WORKER"), controllers.SaveExamination)
		//examination.PUT("", auth.RequireMinRole("WORKER"), controllers.UpdateExamination)
		examination.DELETE("/:id", auth.RequireMinRole("WORKER"), controllers.DeleteExamination)
	}

	if err := r.Run(":8090"); err != nil {
		return
	}
}
