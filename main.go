package main

import (
	"go_restful_mvc/config"
	"go_restful_mvc/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	config.ConnectDB()
	config.Migrate()

	// Set up router
	r := gin.Default()
	routes.RegisterUserRoutes(r)
	routes.RegisterProductRoutes(r)

	// Start the server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
