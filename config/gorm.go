package config

import (
	"fmt"

	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "superuser"
	dbname   = "mygram"
)

func ConnectPostgresGORM() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{})

	return db, nil
}
