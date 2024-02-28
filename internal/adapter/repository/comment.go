package repository

import (
	"ghepa/internal/core/domain"
	"ghepa/internal/core/port"

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

func (c *CommentRepository) Create(req domain.CommentRequest) (*domain.Comment, error) {
	entity := domain.Comment{
		UserID:  req.UserID,
		EventID: req.EventID,
		Comment: req.Comment,
	}

	err := c.db.Create(&entity).Error
	if err != nil {
		return nil, err
	}

	err = c.db.Preload("Event").
		Preload("User").
		Find(&entity).Error
	if err != nil {
		return nil, err
	}

	err = c.db.Model(&entity.Event).
		Association("Comments").
		Append(&entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

func (c *CommentRepository) GetAll() ([]domain.Comment, error) {
	var entity []domain.Comment
	err := c.db.
		Preload("Event").
		Preload("User").
		Find(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}
