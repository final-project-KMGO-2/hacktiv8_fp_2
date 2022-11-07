package repository

import (
	"context"
	"errors"
	"fmt"
	"hacktiv8_fp_2/entity"

	"gorm.io/gorm"
)

type SocmedRepository interface {
	CreateSocmed(ctx context.Context, Socmed entity.SocialMedia) (entity.SocialMedia, error)
	GetSocmedByEmail(ctx context.Context, email string) (entity.SocialMedia, error)
	UpdateSocmed(ctx context.Context, Socmed entity.SocialMedia) (entity.SocialMedia, error)
	DeleteSocmed(ctx context.Context, SocmedID uint64) error
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
	fmt.Print("socmed repo -> ")
	fmt.Printf("%+v\n", socmed)
	tx := sc.connection.Create(&socmed)
	if tx.Error != nil {
		return entity.SocialMedia{}, tx.Error
	}

	return socmed, nil
}

func (sc *socmedConnection) GetSocmedByEmail(ctx context.Context, email string) (entity.SocialMedia, error) {
	return entity.SocialMedia{}, errors.New("hai")
}

func (sc *socmedConnection) UpdateSocmed(ctx context.Context, Socmed entity.SocialMedia) (entity.SocialMedia, error) {
	return entity.SocialMedia{}, errors.New("hai")

}

func (sc *socmedConnection) DeleteSocmed(ctx context.Context, SocmedID uint64) error {
	return errors.New("hai")

}
