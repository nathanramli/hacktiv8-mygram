package models

import "time"

type SocialMedia struct {
	Id             uint `gorm:"primaryKey;autoIncrement"`
	Name           string
	SocialMediaUrl string
	UserId         User `gorm:"foreignKey:u_id;association_foreignKey:id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
