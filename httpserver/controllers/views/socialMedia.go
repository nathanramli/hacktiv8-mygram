package views

import (
	"time"
)

type CreateSocialMedia struct {
	Id             uint   `json:"id"`
	UserId         uint   `json:"user_id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	CreatedAt      time.Time
}

type GetSocialMedia struct {
	Id             uint   `json:"id"`
	UserId         uint   `json:"user_id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           ViewsGetSocialMedia
}

type ViewsGetSocialMedia struct {
	Id             uint   `json:"id"`
	UserName       string `json:"username"`
	SocialMediaUrl string `json:"social_media_url"`
}

type EditSocialMedia struct {
	Id             uint   `json:"id"`
	UserId         uint   `json:"user_id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UpdatedAt      time.Time
}
