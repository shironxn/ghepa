package repository

import (
	"event-planning-app/internal/core/domain"
	"event-planning-app/internal/core/port"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) port.CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (c *CommentRepository) Create(entity domain.Comment) (*domain.Comment, error) {
	err := c.db.Create(&entity).Error
	return &entity, err
}

func (c *CommentRepository) GetAll() ([]domain.Comment, error) {
	var entity []domain.Comment
	err := c.db.Find(&entity).Error
	return entity, err
}
