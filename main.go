package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stefanusong/votify-api/config"
	"github.com/stefanusong/votify-api/models"
	"github.com/stefanusong/votify-api/routes"
	"gorm.io/gorm"
)

func main() {

	r := gin.Default()
	r.Use(cors.Default())

	db := config.OpenDBConnection()

	handleMigration(db, &models.User{}, &models.Category{}, &models.Vote{},
		&models.VoteOption{}, &models.VoteQuestion{}, &models.UserAnswer{})

	routeGroup := r.Group("/api")
	routes.InitRoutes(routeGroup, db)

	r.Run()
}

func handleMigration(db *gorm.DB, dst ...interface{}) {

	fmt.Println("Migrating tables...")
	err := db.AutoMigrate(dst...)
	if err != nil {
		log.Fatal("Failed migrating tables: ", err)
	}
	fmt.Println("Successfuly migrated !")

}
