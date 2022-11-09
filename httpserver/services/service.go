package services

import (
	"context"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/views"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
)

type UserSvc interface {
	Register(ctx context.Context, user *params.Register) *views.Response
	Login(ctx context.Context, user *params.Login) *views.Response
}

type PhotoSvc interface {
	CreatePhoto(ctx context.Context, photo *params.CreatePhoto, UserID int) *views.Response
	GetPhotos(ctx context.Context) *views.Response
	UpdatePhoto(ctx context.Context, photo *params.UpdatePhoto, id int) *views.Response
	GetPhotoByID(ctx context.Context, id int) (*models.Photo, error)
	DeletePhoto(ctx context.Context, id int) *views.Response
}

type CommentSvc interface {
	CreateComment(ctx context.Context, comment *params.CreateComment, UserID int) *views.Response
	GetComments(ctx context.Context) *views.Response
}