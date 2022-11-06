package models

import "time"

type Photo struct {
	Id     	  int       `gorm:"primaryKey;autoIncrement"`
	UserId    int 		`gorm:"foreignKey:Id"`   
	Title     string    
	Caption   string 
	PhotoUrl  string     
	CreatedAt time.Time 
	UpdatedAt time.Time 
}