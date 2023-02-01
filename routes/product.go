package routes

import (
	"github.com/gin-gonic/gin"
	"go-commerce/handlers"
	"go-commerce/middlewares"
	"go-commerce/repositories"
	"go-commerce/services"
	"gorm.io/gorm"
)

func ProductRoute(db *gorm.DB, router *gin.RouterGroup) {

	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	productRouter := router.Group("/products").Use(middlewares.AuthMiddleware())

	productRouter.GET("/", productHandler.FindAll)
}
