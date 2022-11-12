package dto

type PhotoCreateDTO struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
	UserID   uint64 `gorm:"foreignKey" json:"user_id"`
}

type PhotoUpdateDTO struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
	UserID   uint64 `gorm:"foreignKey" json:"user_id"`
}
