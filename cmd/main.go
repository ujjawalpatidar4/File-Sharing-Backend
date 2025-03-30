package main

import (
	"file-sharing-backend/config"
	"file-sharing-backend/routes"
	"file-sharing-backend/workers"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()

	
	// Redis
	config.InitRedis()
	
	// Router
	router := gin.Default()
	routes.SetupRouter(router)
	go workers.DeleteExpiredFiles()

	log.Println("Server is running on port 8080.")
	router.Run(":8080")
}
