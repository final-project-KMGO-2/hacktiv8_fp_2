package dto

type SocialMediaCreateDTO struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required"`
	UserID         uint64
}

type SocialMediaUpdateDTO struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required"`
	UserID         uint64 `json:"user_id"`
}
