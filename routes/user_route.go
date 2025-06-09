package routes

import (
	"go_restful_mvc/controllers"
	"go_restful_mvc/repositories"
	"go_restful_mvc/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userCtrl := controllers.NewUserController(userService)

	r.POST("/register", userCtrl.Register)
	r.POST("/login", userCtrl.Login)
	r.PUT("/user/:id", userCtrl.Update)

	return r
}
