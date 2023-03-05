package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stefanusong/votify-api/handlers"
	"github.com/stefanusong/votify-api/middleware"
	"github.com/stefanusong/votify-api/repositories"
	"github.com/stefanusong/votify-api/services"
	"gorm.io/gorm"
)

func SetUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	userRoute := router.Group("/user")
	userRoute.POST("/", userHandler.Register)
	userRoute.GET("/:username", middleware.AuthMiddleware, userHandler.GetUserByUsername)
}
