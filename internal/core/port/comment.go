package port

import (
	"event-planning-app/internal/core/domain"
	"net/http"
)

type CommentRepository interface {
	Create(entity domain.Event) (*domain.Comment, error)
	GetAll() ([]domain.Comment, error)
}

type CommentService interface {
	Create(entity domain.Event) (*domain.Comment, error)
	GetAll() ([]domain.Comment, error)
}

type CommentHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}
