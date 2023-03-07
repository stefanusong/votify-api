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

	voteRoute := router.Group("/votes")
	voteRoute.POST("/", middleware.AuthMiddleware, voteHandler.CreateVote)
	voteRoute.GET("/:id", middleware.AuthMiddleware, voteHandler.GetVoteByID)
	voteRoute.GET("/public", middleware.AuthMiddleware, voteHandler.GetPublicVotes)
}
