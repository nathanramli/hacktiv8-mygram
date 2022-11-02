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

type SocialMediaSvc interface {
	CreateSocialMedia(ctx context.Context, socialMedia *params.CreateSocialMedia) *views.Response
	GetSocialMedia(ctx context.Context) *views.Response
	UpdateSocialMedia(ctx context.Context, socialMedia *params.UpdateSocialMedia, id uint) *views.Response
}
