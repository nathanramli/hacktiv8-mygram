package services

import (
	"context"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/views"
)

type UserSvc interface {
	Register(ctx context.Context, user *params.Register) *views.Response
	Login(ctx context.Context, user *params.Login) *views.Response
}

type PhotoSvc interface {
	CreatePhoto(ctx context.Context, photo *params.CreatePhoto) *views.Response
	GetPhotos(ctx context.Context) *views.Response
	UpdatePhoto(ctx context.Context, photo *params.UpdatePhoto, id uint) *views.Response
}