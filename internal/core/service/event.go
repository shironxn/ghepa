package service

import (
	"errors"
	"event-planning-app/internal/core/domain"
	"event-planning-app/internal/core/port"
)

type EventService struct {
	repository port.EventRepository
}

func NewEventService(repository port.EventRepository) port.EventService {
	return &EventService{
		repository: repository,
	}
}

func (e *EventService) Create(entity domain.Event) (*domain.Event, error) {
	data, err := e.repository.Create(entity)
	return data, err
}

func (e *EventService) GetAll() ([]domain.Event, error) {
	data, err := e.repository.GetAll()
	return data, err
}

func (e *EventService) GetAllByUser(id uint) ([]domain.Event, error) {
	data, err := e.repository.GetAllByUser(id)
	return data, err
}

func (e *EventService) GetByID(id uint) (*domain.Event, error) {
	data, err := e.repository.GetByID(id)
	return data, err
}

func (e *EventService) Update(entity domain.Event, id uint) (*domain.Event, error) {
	event, err := e.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if event.User.ID != entity.User.ID {
		return nil, errors.New("user does not have permission to perform this action")
	}

	data, err := e.repository.Update(event, entity)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (e *EventService) JoinEvent(entity domain.Participant) (*domain.Participant, error) {
	event, err := e.repository.JoinEvent(entity)
	return event, err
}

func (e *EventService) Delete(entity domain.Event) error {
	event, err := e.repository.GetByID(entity.ID)
	if err != nil {
		return err
	}

	if event.User.ID != entity.User.ID {
		return errors.New("user does not have permission to perform this action")
	}

	err = e.repository.Delete(event)
	if err != nil {
		return err
	}

	return nil
}
