package models

import "time"

type SocialMedia struct {
	Id             uint `gorm:"primaryKey;autoIncrement"`
	UserId         int  `gorm:"foreignKey:Id"`
	Name           string
	SocialMediaUrl string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           User
}
