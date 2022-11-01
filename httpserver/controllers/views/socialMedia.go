package views

import (
	"time"

	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
)

type CreateSocialMedia struct {
	Id             uint        `json:"id"`
	Name           string      `json:"name"`
	SocialMediaUrl string      `json:"social_media_url"`
	UserId         models.User `json:"user_id"`
	CreatedAt      time.Time
}

type EditSocialMedia struct {
	Id             uint        `json:"id"`
	Name           string      `json:"name"`
	SocialMediaUrl string      `json:"social_media_url"`
	UserId         models.User `json:"user_id"`
	UpdatedAt      time.Time
}
