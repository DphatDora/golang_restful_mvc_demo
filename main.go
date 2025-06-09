package main

import (
	"go_restful_mvc/config"
	"go_restful_mvc/routes"
)

func main() {
	config.ConnectDB()
	config.Migrate()
	r := routes.SetupRoutes()
	r.Run(":8080")
}
