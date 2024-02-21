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

func (e *EventService) Create(req domain.EventRequest) (*domain.Event, error) {
	data, err := e.repository.Create(req)
	return data, err
}

func (e *EventService) GetAll() ([]domain.Event, error) {
	data, err := e.repository.GetAll()
	return data, err
}

func (e *EventService) GetAllByUser(req domain.EventRequest) ([]domain.Event, error) {
	data, err := e.repository.GetAllByUser(req.User.ID)
	return data, err
}

func (e *EventService) GetByID(req domain.EventRequest) (*domain.Event, error) {
	data, err := e.repository.GetByID(req.User.ID)
	return data, err
}

func (e *EventService) Update(req domain.EventRequest, claims domain.Claims) (*domain.Event, error) {
	event, err := e.repository.GetByID(req.User.ID)
	if err != nil {
		return nil, err
	}

	if event.User.ID != claims.ID {
		return nil, errors.New("user does not have permission to perform this action")
	}

	data, err := e.repository.Update(event, req)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (e *EventService) Delete(req domain.EventRequest, claims domain.Claims) error {
	event, err := e.repository.GetByID(req.User.ID)
	if err != nil {
		return err
	}

	if event.User.ID != claims.ID {
		return errors.New("user does not have permission to perform this action")
	}

	err = e.repository.Delete(event)
	if err != nil {
		return err
	}

	return nil
}

func (e *EventService) JoinEvent(req domain.ParticipantRequest) (*domain.Participant, error) {
	event, err := e.repository.JoinEvent(req)
	return event, err
}
