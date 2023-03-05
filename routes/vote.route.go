package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stefanusong/votify-api/handlers"
	"github.com/stefanusong/votify-api/middleware"
	"github.com/stefanusong/votify-api/repositories"
	"github.com/stefanusong/votify-api/services"
	"gorm.io/gorm"
)

func SetVoteRoutes(router *gin.RouterGroup, db *gorm.DB) {
	voteRepo := repositories.NewVoteRepository(db)
	voteService := services.NewVoteService(voteRepo)
	voteHandler := handlers.NewVoteHandler(voteService)

	voteRoute := router.Group("/vote")
	voteRoute.POST("/", middleware.AuthMiddleware, voteHandler.CreateVote)
}
