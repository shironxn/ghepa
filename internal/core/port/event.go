package port

import (
	"event-planning-app/internal/core/domain"
	"net/http"
)

type EventRepository interface {
	Create(req domain.Event) (*domain.Event, error)
	GetAll() ([]domain.Event, error)
	GetAllByUser(id uint) ([]domain.Event, error)
	GetByID(id uint) (*domain.Event, error)
	Update(event *domain.Event, req domain.Event) (*domain.Event, error)
	Delete(event *domain.Event) error
}

type EventService interface {
	Create(req domain.Event) (*domain.Event, error)
	GetAll() ([]domain.Event, error)
	GetAllByUser(id uint) ([]domain.Event, error)
	GetByID(id uint) (*domain.Event, error)
	Update(id uint, req domain.Event) (*domain.Event, error)
	Delete(id uint) error
}

type EventHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetAllByUser(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
