package models

import "time"

type User struct {
	Id int `gorm:"primaryKey;autoIncrement"`
	// UserId    int
	Username  string
	Email     string
	Password  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
