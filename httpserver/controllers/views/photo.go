package views

import (
	"time"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
)

type CreatePhoto struct {
	Id       uint    	  `json:"id"`
	Title    string    	  `json:"title"`
	Caption  string 	  `json:"caption"`
	PhotoUrl string 	  `json:"photo_url"`
	UserId 	 models.User  `json:"user_id`
	CreatedAt time.Time 
}

type GetPhotos struct {
	Id       uint    	  `json:"id"`
	Title    string    	  `json:"title"`
	Caption  string 	  `json:"caption"`
	PhotoUrl string 	  `json:"photo_url"`
	UserId 	 models.User `json:"user_id`
	CreatedAt time.Time 
	UpdatedAt time.Time
	User 	 models.User
}
