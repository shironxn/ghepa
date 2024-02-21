package repository

import (
	"event-planning-app/internal/core/domain"
	"event-planning-app/internal/core/port"

	"github.com/charmbracelet/log"
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
		Title:       req.Title,
		Description: req.Description,
		EndDate:     req.EndDate,
		User:        req.User,
	}
	err := e.db.Create(&entity).Error
	return &entity, err
}

func (e *EventRepository) GetAll() ([]domain.Event, error) {
	var entity []domain.Event
	if err := e.db.Preload("User").
		Preload("Comments").Preload("Comments.User").
		Preload("Participants.User").
		Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (e *EventRepository) GetAllByUser(userID uint) ([]domain.Event, error) {
	var entity []domain.Event
	if err := e.db.Preload("User").
		Preload("Comments").Preload("Comments.User").
		Preload("Participants.User").
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
		Preload("Participants.User").
		First(&entity, eventID).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (e *EventRepository) Update(entity *domain.Event, req domain.EventRequest) (*domain.Event, error) {
	entityUpdate := domain.Event{
		Title:       req.Title,
		Description: req.Description,
		EndDate:     req.EndDate,
		User:        req.User,
	}
	err := e.db.Model(entity).Updates(entityUpdate).Error
	return &entityUpdate, err
}

func (e *EventRepository) Delete(entity *domain.Event) error {
	err := e.db.Delete(&entity).Error
	return err
}

func (e *EventRepository) JoinEvent(req domain.ParticipantRequest) (*domain.Participant, error) {
	entity := domain.Participant{
		UserID:  req.UserID,
		EventID: req.EventID,
	}

	err := e.db.Create(&entity).Error
	if err != nil {
		return nil, err
	}

	err = e.db.Model(&entity.Event).Association("Participants").Append(&entity)
	if err != nil {
		log.Info("tes2")
		return nil, err
	}

	return &entity, nil
}
