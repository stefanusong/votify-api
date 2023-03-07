package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stefanusong/votify-api/handlers"
	"github.com/stefanusong/votify-api/middleware"
	"github.com/stefanusong/votify-api/repositories"
	"github.com/stefanusong/votify-api/services"
	"gorm.io/gorm"
)

func SetCategoryRoutes(router *gin.RouterGroup, db *gorm.DB) {
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	categoryRoute := router.Group("/categories")
	categoryRoute.POST("/", middleware.AuthMiddleware, categoryHandler.CreateCategory)
	categoryRoute.GET("/", middleware.AuthMiddleware, categoryHandler.GetAllCategories)
	categoryRoute.GET("/:id", middleware.AuthMiddleware, categoryHandler.GetCategoryByID)
	categoryRoute.PUT("/:id", middleware.AuthMiddleware, categoryHandler.UpdateCategoryByID)
	categoryRoute.DELETE("/:id", middleware.AuthMiddleware, categoryHandler.DeleteCategoryByID)
}
