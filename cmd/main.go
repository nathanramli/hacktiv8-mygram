package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nathanramli/hacktiv8-mygram/config"
	"github.com/nathanramli/hacktiv8-mygram/httpserver"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/gorm"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/services"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("cannot  load env")
		return
	}
}

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

	app := httpserver.NewRouter(router, userHandler)
	app.Start(":" + os.Getenv("PORT"))
}
