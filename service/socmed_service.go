package service

import (
	"context"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/repository"
)

type SocmedService interface {
	GetSocmedInfo(ctx context.Context) (entity.SocialMedia, error)
	AddNewSocmed(ctx context.Context, email string) (entity.SocialMedia, error)
	DeleteSocmed(ctx context.Context, email string) (entity.SocialMedia, error)
	UpdateSocmed(ctx context.Context, email string) (entity.SocialMedia, error)
}

type socmedService struct {
	socmedRepository repository.SocmedRepository
}

func NewSocmedService(sr repository.SocmedRepository) SocmedService {
	return &socmedService{
		socmedRepository: sr,
	}
}

func (sr *socmedService) GetSocmedInfo(ctx context.Context) (entity.SocialMedia, error) {}
func (sr *socmedService) AddNewSocmed(ctx context.Context, email string) (entity.SocialMedia, error) {}
func (sr *socmedService) DeleteSocmed(ctx context.Context, email string) (entity.SocialMedia, error) {}
func (sr *socmedService) UpdateSocmed(ctx context.Context, email string) (entity.SocialMedia, error) {}






