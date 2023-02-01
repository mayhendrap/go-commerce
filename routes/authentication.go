package routes

import (
	"github.com/gin-gonic/gin"
	"go-commerce/handlers"
	"go-commerce/repositories"
	"go-commerce/services"
	"gorm.io/gorm"
)

func AuthenticationRoute(db *gorm.DB, router *gin.RouterGroup) {

	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthenticationService(userRepository)
	authHandler := handlers.NewAuthenticationHandler(authService)

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
}
