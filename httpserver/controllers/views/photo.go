package views

import (
	"time"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
)

type CreatePhoto struct {
	Id       int    	  `json:"id"`
	Title    string    	  `json:"title"`
	Caption  string 	  `json:"caption"`
	PhotoUrl string 	  `json:"photo_url"`
	UserId 	 int  		  `json:"user_id"`
	CreatedAt time.Time   `json:"created_at"`
}

type GetPhotos struct {
	Id       int    	  `json:"id"`
	Title    string    	  `json:"title"`
	Caption  string 	  `json:"caption"`
	PhotoUrl string 	  `json:"photo_url"`
	UserId 	 int 		  `json:"user_id"`
	CreatedAt time.Time 
	UpdatedAt time.Time
	User 	 models.User  `json:"user"`
}

type UpdatePhoto struct {
	Id       int    	  `json:"id"`
	Title    string    	  `json:"title"`
	Caption  string 	  `json:"caption"`
	PhotoUrl string 	  `json:"photo_url"`
	UserId 	 int  `json:"user_id"`
	UpdatedAt time.Time 
}
