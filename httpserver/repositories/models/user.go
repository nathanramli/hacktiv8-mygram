package models

import "time"

type User struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	Username  string
	Email     string
	Password  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserSocialMedia struct {
	UserId   int `gorm:"foreignKey:Id"`
	Username string
}
