package views

type Register struct {
	Id       int    `json:"id"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Login struct {
	Token string `json:"token"`
}
