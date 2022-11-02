package models

import "time"

type Photo struct {
	Id     	  uint       `gorm:"primaryKey;autoIncrement"`
	Title     string    
	Caption   string 
	PhotoUrl  string    
	UserId    User 		`gorm:"foreignKey:Id"`    
	CreatedAt time.Time 
	UpdatedAt time.Time 
}