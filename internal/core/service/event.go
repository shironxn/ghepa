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

func (e *EventService) Create(req domain.Event) (*domain.Event, error) {
	data, err := e.repository.Create(req)
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

func (e *EventService) Update(req domain.Event, id uint) (*domain.Event, error) {
	event, err := e.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	if event.User.ID != req.User.ID {
		return nil, errors.New("user does not have permission to perform this action")
	}

	data, err := e.repository.Update(event, req)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (e *EventService) Delete(req domain.Event) error {
	event, err := e.repository.GetByID(req.ID)
	if err != nil {
		return err
	}

	if event.User.ID != req.User.ID {
		return errors.New("user does not have permission to perform this action")
	}

	err = e.repository.Delete(event)
	if err != nil {
		return err
	}

	return nil
}
