package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stefanusong/votify-api/handlers"
	"github.com/stefanusong/votify-api/middleware"
	"github.com/stefanusong/votify-api/repositories"
	"github.com/stefanusong/votify-api/services"
	"gorm.io/gorm"
)

func SetAuthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	authRoute := router.Group("/auth")
	authRoute.POST("/login", authHandler.Login)
	authRoute.GET("/check", middleware.AuthMiddleware, authHandler.AuthCheck)
}
