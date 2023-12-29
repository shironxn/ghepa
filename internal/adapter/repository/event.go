package repository

import (
	"event-planning-app/internal/core/domain"
	"event-planning-app/internal/core/port"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) port.EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (e *EventRepository) Create(entity domain.Event) (*domain.Event, error) {
	err := e.db.Create(&entity).Error
	return &entity, err
}

func (e *EventRepository) GetAll() ([]domain.Event, error) {
	var entity []domain.Event
	if err := e.db.Preload("User").
		Preload("Comments").Preload("Comments.User").
		Preload("Participant.User").
		Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (e *EventRepository) GetAllByUser(userID uint) ([]domain.Event, error) {
	var entity []domain.Event
	if err := e.db.Preload("User").
		Preload("Comments").Preload("Comments.User").
		Preload("Participant.User").
		Where("user_id = ?", userID).
		Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (e *EventRepository) GetByID(eventID uint) (*domain.Event, error) {
	var entity domain.Event
	if err := e.db.Preload("User").
		Preload("Comments").Preload("Comments.User").
		Preload("Participant.User").
		First(&entity, eventID).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (e *EventRepository) Update(entity *domain.Event, entityUpdate domain.Event) (*domain.Event, error) {
	err := e.db.Model(&entity).Updates(entityUpdate).Error
	return entity, err
}

func (e *EventRepository) Delete(entity *domain.Event) error {
	err := e.db.Delete(&entity).Error
	return err
}
