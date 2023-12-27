package repository

import (
	"errors"
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

func (e *EventRepository) Create(event domain.Event) (*domain.Event, error) {
	err := e.db.Create(&event).Error
	return &event, err
}

func (e *EventRepository) GetAll() ([]domain.Event, error) {
	var events []domain.Event
	if err := e.db.Preload("User").
		Preload("Comments").Preload("Comments.User").
		Preload("Participant.User").
		Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (e *EventRepository) GetAllByUser(userID uint) ([]domain.Event, error) {
	var events []domain.Event
	if err := e.db.Preload("User").
		Preload("Comments").Preload("Comments.User").
		Preload("Participant.User").
		Where("user_id = ?", userID).
		Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (e *EventRepository) GetByID(eventID uint) (*domain.Event, error) {
	var event domain.Event
	if err := e.db.Preload("User").
		Preload("Comments").Preload("Comments.User").
		Preload("Participant.User").
		First(&event, eventID).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (e *EventRepository) Update(event *domain.Event, updateData domain.Event) (*domain.Event, error) {
	result := e.db.Model(event).Where("user_id = ?", updateData.User.ID).Updates(updateData)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no rows affected by update")
	}

	return event, nil
}

func (e *EventRepository) Delete(event *domain.Event) error {
	err := e.db.Delete(event).Error
	return err
}
