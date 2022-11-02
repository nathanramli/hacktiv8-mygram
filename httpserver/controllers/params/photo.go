package params

type CreatePhoto struct {
	Title      string `json:"title"`
	Caption    string `json:"caption"`
	PhotoUrl   string `json:"photo_url"`
}