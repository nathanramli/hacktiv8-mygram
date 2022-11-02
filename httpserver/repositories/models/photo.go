package models

import "time"

type Photo struct {
	Id     	  uint       `gorm:"primaryKey;autoIncrement"`
	UserId    uint 		`gorm:"foreignKey:Id"`   
	Title     string    
	Caption   string 
	PhotoUrl  string     
	CreatedAt time.Time 
	UpdatedAt time.Time 
}