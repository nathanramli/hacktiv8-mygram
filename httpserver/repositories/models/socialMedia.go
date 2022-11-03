package models

import "time"

type SocialMedia struct {
	Id             uint `gorm:"primaryKey;autoIncrement"`
	UserId 	   		int
	User         	User `gorm:"foreignKey:UserId"`
	Name           string
	SocialMediaUrl string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
