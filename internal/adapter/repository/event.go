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

func (e *EventRepository) Create(req domain.Event) (*domain.Event, error) {
	err := e.db.Create(&req).Error
	if err != nil {
		return nil, err
	}

	err = e.db.Preload("User").Find(&req).Error
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (e *EventRepository) GetAll() ([]domain.Event, error) {
	var events []domain.Event
	err := e.db.Preload("User").Preload("Comments").Preload("Comments.User").Preload("Participant.User").Find(&events).Error
	return events, err
}

func (e *EventRepository) GetByID(id uint) (*domain.Event, error) {
	var event domain.Event
	err := e.db.Preload("User").Preload("Comments").Preload("Comments.User").Preload("Participant.User").First(&event, id).Error
	return &event, err
}

func (e *EventRepository) Update(event *domain.Event, req domain.Event) (*domain.Event, error) {
	err := e.db.Model(event).Updates(req).Error
	return event, err
}

func (e *EventRepository) Delete(event *domain.Event) error {
	err := e.db.Delete(event).Error
	return err
}
