package service

import (
	"context"
	"errors"
	"fmt"
	"hacktiv8_fp_2/dto"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/repository"

	"github.com/mashingan/smapping"
)

type SocmedService interface {
	GetSocmedInfo(ctx context.Context) (entity.SocialMedia, error)
	AddNewSocmed(ctx context.Context, socmedDTO dto.SocialMediaCreateDTO) (entity.SocialMedia, error)
	DeleteSocmed(ctx context.Context, email string) (entity.SocialMedia, error)
	UpdateSocmed(ctx context.Context, email string) (entity.SocialMedia, error)
}

type socmedService struct {
	socmedRepository repository.SocmedRepository
	userRepository   repository.UserRepository
}

func NewSocmedService(sr repository.SocmedRepository, ur repository.UserRepository) SocmedService {
	return &socmedService{
		socmedRepository: sr,
		userRepository:   ur,
	}
}

func (sr *socmedService) GetSocmedInfo(ctx context.Context) (entity.SocialMedia, error) {
	return entity.SocialMedia{}, errors.New("hai")

}
func (sr *socmedService) AddNewSocmed(ctx context.Context, socmedDTO dto.SocialMediaCreateDTO) (entity.SocialMedia, error) {
	socmed := entity.SocialMedia{}
	err := smapping.FillStruct(&socmed, smapping.MapFields(&socmedDTO))
	if err != nil {
		return entity.SocialMedia{}, err
	}

	fmt.Print("ini scmed -> ")
	fmt.Printf("%+v\n", socmedDTO)

	res, err := sr.socmedRepository.CreateSocmed(ctx, socmed)

	user, err := sr.userRepository.GetUserById(ctx, socmed.UserID)
	if err != nil {
		return entity.SocialMedia{}, err
	}
	fmt.Print("user -> ");
	fmt.Printf("%+v\n", user)
	fmt.Println("socmed -> ", socmed);

	err = smapping.FillStruct(&res, smapping.MapFields(&user))
	if err != nil {
		return entity.SocialMedia{}, err
	}
	
	return res, nil
}
func (sr *socmedService) DeleteSocmed(ctx context.Context, email string) (entity.SocialMedia, error) {
	return entity.SocialMedia{}, errors.New("hai")

}
func (sr *socmedService) UpdateSocmed(ctx context.Context, email string) (entity.SocialMedia, error) {
	return entity.SocialMedia{}, errors.New("hai")

}
