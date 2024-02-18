package service

import (
	"event-planning-app/internal/core/domain"
	"event-planning-app/internal/core/port"
)

type CommentService struct {
	repository port.CommentRepository
}

func NewCommentService(repository port.CommentRepository) port.CommentService {
	return &CommentService{
		repository: repository,
	}
}

func (c *CommentService) Create(entity domain.Comment) (*domain.Comment, error) {
	data, err := c.repository.Create(entity)
	return data, err
}

func (c *CommentService) GetAll() ([]domain.Comment, error) {
	data, err := c.repository.GetAll()
	return data, err
}
