package params

type Register struct {
	Age      int    `json:"age"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username"`
}
