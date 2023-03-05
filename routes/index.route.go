package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	SetAuthRoutes(rg, db)
	SetUserRoutes(rg, db)
	SetVoteRoutes(rg, db)
}
