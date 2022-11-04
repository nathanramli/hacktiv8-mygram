package views

import (
	"time"

	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
)

type CreateSocialMedia struct {
	Id             uint   `json:"id"`
	UserId         uint   `json:"user_id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	CreatedAt      time.Time
	// UserId int  `json:"user_id" gorm:"foreignKey:UserId"`
	// UserId 		models.User
}

type GetSocialMedia struct {
	Id             uint   `json:"id"`
	UserId         uint   `json:"user_id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           models.User `json:"user"`
	// UserId         int    `json:"user_id" gorm:"foreignKey:UserId"`
}

type EditSocialMedia struct {
	Id             uint   `json:"id"`
	UserId         uint   `json:"user_id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UpdatedAt      time.Time
	// UserId         int    `json:"user_id" gorm:"foreignKey:UserId"`
}

// type DeleteSocialMedia struct {
// 	Message string `json:"message"`
// }
