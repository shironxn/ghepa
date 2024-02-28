package port

import (
	"ghepa/internal/core/domain"
	"net/http"
)

type CommentRepository interface {
	Create(req domain.CommentRequest) (*domain.Comment, error)
	GetAll() ([]domain.Comment, error)
}

type CommentService interface {
	Create(req domain.CommentRequest) (*domain.Comment, error)
	GetAll() ([]domain.Comment, error)
}

type CommentHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}
