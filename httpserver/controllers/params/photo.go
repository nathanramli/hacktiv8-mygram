package params

type CreatePhoto struct {
	Title      string `json:"title" validate:"required"`
	Caption    string `json:"caption"`
	PhotoUrl   string `json:"photo_url" validate:"required"`
}

type UpdatePhoto struct {
	Title      string `json:"title"`
	Caption    string `json:"caption"`
	PhotoUrl   string `json:"photo_url"`
}