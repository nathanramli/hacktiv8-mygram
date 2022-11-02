package params

type CreateSocialMedia struct {
	Name string `json:"name" validate:"required"`
	// SocialMediaUrl string `json:"social_media_url" validate:"required,social_media_url"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}

type UpdateSocialMedia struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}
