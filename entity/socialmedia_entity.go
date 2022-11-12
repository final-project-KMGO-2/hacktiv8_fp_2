package entity

type SocialMedia struct {
	ID             uint64 `gorm:"primaryKey" json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
	UserID         uint64 `gorm:"foreignKey" json:"user_id"`
	User           User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
}
