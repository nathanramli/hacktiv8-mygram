package models

import 
	"time"
	

type Comment struct {
	Id     	  int       `gorm:"primaryKey;autoIncrement"`
	UserId    int 		`gorm:"foreignKey:Id"`   
	PhotoId   int       `gorm:"foreignKey:Id"`
	Message     string    
	CreatedAt time.Time 
	UpdatedAt time.Time 
}