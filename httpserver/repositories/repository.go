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

type PhotoRepo interface {
	CreatePhoto(ctx context.Context, photo *models.Photo) error
	GetPhotos(ctx context.Context) ([]models.Photo, error)
	UpdatePhoto(ctx context.Context, photo *models.Photo, id int) error
	FindPhotoByID(ctx context.Context, id int) (*models.Photo, error)
	DeletePhoto(ctx context.Context, id int) error
}

type CommentRepo interface {
	CreateComment(ctx context.Context, comment *models.Comment) error
	GetComments(ctx context.Context) ([]models.Comment, error)
	FindCommentByID(ctx context.Context, id int) (*models.Comment, error)
	UpdateComment(ctx context.Context, comment *models.Comment, id int) error
	DeleteComment(ctx context.Context, id int) error
}
