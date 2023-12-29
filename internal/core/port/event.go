package port

import (
	"event-planning-app/internal/core/domain"
	"net/http"
)

type EventRepository interface {
	Create(entity domain.Event) (*domain.Event, error)
	GetAll() ([]domain.Event, error)
	GetAllByUser(id uint) ([]domain.Event, error)
	GetByID(id uint) (*domain.Event, error)
	Update(entity *domain.Event, entityUpdate domain.Event) (*domain.Event, error)
	Delete(entity *domain.Event) error
}

type EventService interface {
	Create(entity domain.Event) (*domain.Event, error)
	GetAll() ([]domain.Event, error)
	GetAllByUser(id uint) ([]domain.Event, error)
	GetByID(id uint) (*domain.Event, error)
	Update(entity domain.Event, id uint) (*domain.Event, error)
	Delete(entity domain.Event) error
}

type EventHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetAllByUser(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
