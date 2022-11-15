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
	UpdateUser(ctx context.Context, id int, user *params.UpdateUser) *views.Response
	DeleteUser(ctx context.Context, id int) *views.Response
}

type SocialMediaSvc interface {
	CreateSocialMedia(ctx context.Context, socialMedia *params.CreateSocialMedia, UserID uint) *views.Response
	GetSocialMedia(ctx context.Context, user *params.Register) *views.Response
	UpdateSocialMedia(ctx context.Context, socialMedia *params.UpdateSocialMedia, id uint) *views.Response
	DeleteSocialMedia(ctx context.Context, id uint) (*views.Response, error)
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
	UpdateComment(ctx context.Context, comment *params.UpdateComment, id int) *views.Response
	GetCommentByID(ctx context.Context, id int) (*models.Comment, error)
	DeleteComment(ctx context.Context, id int) *views.Response
}
