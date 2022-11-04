package services

import (
	"context"

	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/views"
)

type UserSvc interface {
	Register(ctx context.Context, user *params.Register) *views.Response
	Login(ctx context.Context, user *params.Login) *views.Response
	UpdateUser(ctx context.Context, id int, user *params.UpdateUser) *views.Response
	DeleteUser(ctx context.Context, id int) *views.Response
}

type SocialMediaSvc interface {
	CreateSocialMedia(ctx context.Context, socialMedia *params.CreateSocialMedia, UserID uint) *views.Response
	GetSocialMedia(ctx context.Context) *views.Response
	UpdateSocialMedia(ctx context.Context, socialMedia *params.UpdateSocialMedia, id uint) *views.Response
	DeleteSocialMedia(ctx context.Context, id uint) (*views.Response, error)
}
