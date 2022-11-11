package repository

import (
	"context"
	"fmt"
	"hacktiv8_fp_2/entity"

	"gorm.io/gorm"
)

type SocmedRepository interface {
	CreateSocmed(ctx context.Context, Socmed entity.SocialMedia) (entity.SocialMedia, error)
	GetSocmeds(ctx context.Context) ([]entity.SocialMedia, error)
	UpdateSocmed(ctx context.Context, socmed entity.SocialMedia) (entity.SocialMedia, error)
	DeleteSocmed(ctx context.Context, id uint) error
}

type socmedConnection struct {
	connection *gorm.DB
}

func NewSocmedRepository(db *gorm.DB) SocmedRepository {
	return &socmedConnection{
		connection: db,
	}
}

func (sc *socmedConnection) CreateSocmed(ctx context.Context, socmed entity.SocialMedia) (entity.SocialMedia, error) {
	tx := sc.connection.Preload("User").Create(&socmed)
	fmt.Print("socmed repo -> ")
	fmt.Printf("%+v\n", socmed)
	if tx.Error != nil {
		return entity.SocialMedia{}, tx.Error
	}

	return socmed, nil
}

func (sc *socmedConnection) GetSocmeds(ctx context.Context) ([]entity.SocialMedia, error) {
	var socmed []entity.SocialMedia
	tx := sc.connection.Preload("Users").Find(&socmed)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return socmed, nil
}

func (sc *socmedConnection) UpdateSocmed(ctx context.Context, socmed entity.SocialMedia) (entity.SocialMedia, error) {
	tx := sc.connection.Save(&socmed)
	if tx.Error != nil {
		return entity.SocialMedia{}, tx.Error
	}

	return socmed, nil
}

func (sc *socmedConnection) DeleteSocmed(ctx context.Context, id uint) error {
	tx := sc.connection.Delete(&entity.SocialMedia{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil

}
