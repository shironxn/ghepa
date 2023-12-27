package service

import (
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
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e *EventService) GetAll() ([]domain.Event, error) {
	data, err := e.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e *EventService) GetAllByUser(id uint) ([]domain.Event, error) {
	data, err := e.repository.GetAllByUser(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e *EventService) GetByID(id uint) (*domain.Event, error) {
	data, err := e.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e *EventService) Update(id uint, req domain.Event) (*domain.Event, error) {
	event, err := e.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	data, err := e.repository.Update(event, req)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (e *EventService) Delete(id uint) error {
	event, err := e.repository.GetByID(id)
	if err != nil {
		return err
	}

	err = e.repository.Delete(event)
	if err != nil {
		return err
	}

	return nil
}
