package params

type CreateComment struct {
	Message string `json:"message" validate:"required"`
	PhotoId int    `json:"photo_id" validate:"required"`
}

type UpdateComment struct {
	Message string `json:"message" validate:"required"`
}
