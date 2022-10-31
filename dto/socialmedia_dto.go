package dto

type SocialMediaCreateDTO struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         uint64 `json:"user_id"`
}

type SocialMediaUpdateDTO struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         uint64 `json:"user_id"`
}
