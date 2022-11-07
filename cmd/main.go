package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanramli/hacktiv8-mygram/config"
	"github.com/nathanramli/hacktiv8-mygram/httpserver"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/gorm"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/services"
)

func main() {
	db, err := config.ConnectPostgresGORM()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	config.GenerateJwtSignature()

	userRepo := gorm.NewUserRepo(db)
	userSvc := services.NewUserSvc(userRepo)
	userHandler := controllers.NewUserController(userSvc)

	photoRepo := gorm.NewPhotoRepo(db)
	photoSvc := services.NewPhotoSvc(photoRepo, userRepo)
	photoHandler := controllers.NewPhotoController(photoSvc)

	app := httpserver.NewRouter(router, userHandler, photoHandler)
	app.Start(":8000")
}
