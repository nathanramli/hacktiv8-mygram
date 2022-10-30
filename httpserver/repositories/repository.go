package repositories

import (
	"context"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	FindUserByID(ctx context.Context, id int) (*models.User, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	DeleteUser(ctx context.Context, id int) error
}
