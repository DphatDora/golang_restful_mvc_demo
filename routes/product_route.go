package routes

import (
	"go_restful_mvc/controllers"
	"go_restful_mvc/repositories"
	"go_restful_mvc/services"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	productRepo := repositories.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productCtrl := controllers.NewProductController(productService)

	product := r.Group("/products")
	{
		product.GET("/", productCtrl.FindAll)
		product.GET("/:id", productCtrl.FindByID)
		product.POST("/", productCtrl.Create)
		product.PUT("/:id", productCtrl.Update)
		product.DELETE("/:id", productCtrl.Delete)
	}
}
