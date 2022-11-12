package dto

type CommentCreateDTO struct {
	UserID  uint64 `json:"user_id"`
	PhotoID uint64 `json:"photo_id"`
	Message string `json:"message"`
}

type CommentUpdateDTO struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	UserID  uint64 `json:"user_id"`
	Message string `json:"message"`
}
