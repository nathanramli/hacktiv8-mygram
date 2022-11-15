package views

import (
	"time"
	// "github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
)

type CreateComment struct {
	Id       int    	  `json:"id"`
	Message  string       `json:"message"`
	PhotoId  int 	      `json:"photo_id"`
	UserId 	 int  		  `json:"user_id"`
	CreatedAt time.Time   `json:"created_at"`
}

type GetComments struct {
	Id       int    	  `json:"id"`
	Message    string    	  `json:"message"`
	PhotoId  int 	  `json:"photo_id"`
	UserId 	 int 		  `json:"user_id"`
	UpdatedAt time.Time	  `json:"updated_at"`
	CreatedAt time.Time   `json:"created_at"`
	User 	 UserGetComment  `json:"User"`
	Photo    PhotoGetComment `json:"Photo"`
}

type UserGetComment struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoGetComment struct {
	Id       int    	  `json:"id"`
	Title    string    	  `json:"title"`
	Caption  string 	  `json:"caption"`
	PhotoUrl string 	  `json:"photo_url"`
	UserId 	 int  		  `json:"user_id"`
}

type UpdateComment struct {
	Id       int    	  `json:"id"`
	Message  string       `json:"message"`
	PhotoId  int 	      `json:"photo_id"`
	UserId 	 int 		  `json:"user_id"`
	UpdatedAt time.Time	  `json:"updated_at"`
}

