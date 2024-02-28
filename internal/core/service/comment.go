package service

import (
	"ghepa/internal/core/domain"
	"ghepa/internal/core/port"
)

type CommentService struct {
	repository port.CommentRepository
}

func NewCommentService(repository port.CommentRepository) port.CommentService {
	return &CommentService{
		repository: repository,
	}
}

func (c *CommentService) Create(req domain.CommentRequest) (*domain.Comment, error) {
	data, err := c.repository.Create(req)
	return data, err
}

func (c *CommentService) GetAll() ([]domain.Comment, error) {
	data, err := c.repository.GetAll()
	return data, err
}
