package repository

import (
	"errors"
	"ghepa/internal/core/domain"
	"ghepa/internal/core/port"

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

func (e *EventRepository) Create(req domain.EventRequest) (*domain.Event, error) {
	entity := domain.Event{
		Name:        req.Name,
		Description: req.Description,
		EndDate:     req.EndDate,
		User:        req.User,
	}
	err := e.db.Create(&entity).Error
	return &entity, err
}

func (e *EventRepository) GetAll() ([]domain.Event, error) {
	var entities []domain.Event
	if err := e.db.
		Preload("User").
		Preload("Comments").
		Preload("Comments.User").
		Preload("Participants.User").
		Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (e *EventRepository) GetAllByUser(id uint) ([]domain.Event, error) {
	var entity []domain.Event
	if err := e.db.
		Preload("User").
		Preload("Comments").
		Preload("Comments.User").
		Preload("Participants.User").
		Where("user_id = ?", id).
		Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (e *EventRepository) GetByID(id uint) (*domain.Event, error) {
	var entity domain.Event
	if err := e.db.
		Preload("User").
		Preload("Comments").
		Preload("Comments.User").
		Preload("Participants.User").
		First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (e *EventRepository) Update(entity *domain.Event, req domain.EventRequest) (*domain.Event, error) {
	entityUpdate := domain.Event{
		Name:        req.Name,
		Description: req.Description,
		EndDate:     req.EndDate,
		User:        entity.User,
	}
	err := e.db.Model(entity).Updates(entityUpdate).Error
	if err != nil {
		return nil, err
	}

	err = e.db.First(&entity, entity.ID).Error
	if err != nil {
		return nil, err
	}

	return entity, err
}

func (e *EventRepository) Delete(entity *domain.Event) error {
	err := e.db.Delete(&entity).Error
	return err
}

func (e *EventRepository) JoinEvent(req domain.ParticipantRequest) (*domain.Participant, error) {
	var event domain.Event
	err := e.db.First(&event, req.EventID).Error
	if err != nil {
		return nil, err
	}

	var existingParticipant domain.Participant
	err = e.db.Where("user_id = ? AND event_id = ?", req.UserID, req.EventID).First(&existingParticipant).Error
	if err == nil {
		return nil, errors.New("user already joined the event")
	}

	entity := domain.Participant{
		UserID:  req.UserID,
		EventID: req.EventID,
	}
	err = e.db.Create(&entity).Error
	if err != nil {
		return nil, err
	}

	err = e.db.Preload("Event").Preload("User").First(&entity, entity.ID).Error
	if err != nil {
		return nil, err
	}

	err = e.db.Model(&event).Association("Participants").Append(&entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}
