package models

import (
	"time"

	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/views"
)

type SocialMedia struct {
	Id             uint `gorm:"primaryKey;autoIncrement"`
	UserId         int  `gorm:"foreignKey:Id"`
	Name           string
	SocialMediaUrl string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           views.ViewsGetSocialMedia
}
