package views

import (
	"time"

	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
)

type CreateSocialMedia struct {
	Id             uint        `json:"id"`
	UserId         models.User `json:"user_id" gorm:"foreignKey:Id"`
	Name           string      `json:"name"`
	SocialMediaUrl string      `json:"social_media_url"`
	CreatedAt      time.Time
}

type GetSocialMedia struct {
	Id             uint        `json:"id"`
	UserId         models.User `json:"user_id" gorm:"foreignKey:Id"`
	Name           string      `json:"name"`
	SocialMediaUrl string      `json:"social_media_url"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           models.User
}

type EditSocialMedia struct {
	Id             uint        `json:"id"`
	UserId         models.User `json:"user_id" gorm:"foreignKey:Id"`
	Name           string      `json:"name"`
	SocialMediaUrl string      `json:"social_media_url"`
	UpdatedAt      time.Time
}
