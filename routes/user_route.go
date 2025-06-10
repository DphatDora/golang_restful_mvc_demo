package routes

import (
	"go_restful_mvc/controllers"
	"go_restful_mvc/repositories"
	"go_restful_mvc/services"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userCtrl := controllers.NewUserController(userService)

	auth := r.Group("/auth")
	{
		auth.POST("/register", userCtrl.Register)
		auth.POST("/login", userCtrl.Login)
		auth.PUT("/user/:id", userCtrl.Update)
	}
}
