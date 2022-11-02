package views

import "time"

type Register struct {
	Id       int    `json:"id"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Login struct {
	Token string `json:"token"`
}

type UpdateUser struct {
	Id        int       `json:"id"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	UpdatedAt time.Time `json:"updated_at"`
}
