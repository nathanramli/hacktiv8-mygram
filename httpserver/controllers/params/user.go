package params

type Register struct {
	Age      int    `json:"age" validate:"required,gt=8"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Username string `json:"username" validate:"required"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateUser struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
}
