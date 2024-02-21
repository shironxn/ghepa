package port

import (
	"event-planning-app/internal/core/domain"
	"net/http"
)

type EventRepository interface {
	Create(req domain.EventRequest) (*domain.Event, error)
	GetAll() ([]domain.Event, error)
	GetAllByUser(id uint) ([]domain.Event, error)
	GetByID(id uint) (*domain.Event, error)
	Update(entity *domain.Event, req domain.EventRequest) (*domain.Event, error)
	Delete(entity *domain.Event) error
	JoinEvent(req domain.ParticipantRequest) (*domain.Participant, error)
}

type EventService interface {
	Create(req domain.EventRequest) (*domain.Event, error)
	GetAll() ([]domain.Event, error)
	GetAllByUser(req domain.EventRequest) ([]domain.Event, error)
	GetByID(req domain.EventRequest) (*domain.Event, error)
	Update(req domain.EventRequest, claims domain.Claims) (*domain.Event, error)
	Delete(req domain.EventRequest, claims domain.Claims) error
	JoinEvent(req domain.ParticipantRequest) (*domain.Participant, error)
}

type EventHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetAllByUser(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	JoinEvent(w http.ResponseWriter, r *http.Request)
}
