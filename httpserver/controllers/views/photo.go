package views

import (
	"time"
	// "github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
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
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time	  `json:"updated_at"`
	User 	 UserGetPhoto  `json:"User"`
}

type UserGetPhoto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdatePhoto struct {
	Id       int    	  `json:"id"`
	Title    string    	  `json:"title"`
	Caption  string 	  `json:"caption"`
	PhotoUrl string 	  `json:"photo_url"`
	UserId 	 int  `json:"user_id"`
	UpdatedAt time.Time 
}
