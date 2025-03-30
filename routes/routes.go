package routes

import (
	"file-sharing-backend/handlers"
	"file-sharing-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	// Authentication Routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", handlers.Register)
		authRoutes.POST("/login", handlers.Login)
	}

	// File Management Routes (with middleware)
	fileRoutes := router.Group("/files")
	fileRoutes.Use(middleware.AuthMiddleware())
	{
		fileRoutes.POST("/upload", handlers.FileUploadHandler)
		fileRoutes.GET("/list", handlers.ListFiles)
		fileRoutes.GET("/download/:filename", handlers.DownloadFile)
		fileRoutes.DELETE("/delete/:filename", handlers.DeleteFile)

		fileRoutes.GET("/search", handlers.SearchFiles)

		fileRoutes.DELETE("/delete-expired", handlers.TriggerDeletion)
	}
}
