package service

import (
	"context"
	"hacktiv8_fp_2/dto"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/repository"
)

type CommentService interface {
	CreateComment(ctx context.Context, commentDTO dto.CommentCreateDTO) (entity.Comment, error)
	GetComment(ctx context.Context) ([]entity.Comment, error)
	UpdateCommentByID(ctx context.Context, commentDTO dto.CommentUpdateDTO, id string) (entity.Comment, error)
	DeleteCommentByID(ctx context.Context, id string) error
}

type commentService struct {
	CommentRepository repository.CommentConnection
}

func NewCommentService(cr repository.CommentRepository) CommentService {
	return &commentService{
		CommentRepository: cr,
	}
}

func (s *commentService) CreateComment(ctx context.Context, commentDTO dto.CommentCreateDTO) entity.Comment
