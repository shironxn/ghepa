package port

import (
	"event-planning-app/internal/core/domain"
	"net/http"
)

type UserRepository interface {
	Create(entity domain.User) (*domain.User, error)
	GetAll() ([]domain.User, error)
	GetByID(id uint) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Update(entity *domain.User, entityUpdate domain.User) (*domain.User, error)
	Delete(entity *domain.User) error
}

type UserService interface {
	Create(entity domain.RegisterRequest) (*domain.User, error)
	Login(entity domain.LoginRequest) (*domain.User, error)
	GetAll() ([]domain.User, error)
	GetByID(id uint) (*domain.User, error)
	Update(entity domain.User, id uint) (*domain.User, error)
	Delete(entity domain.User) error
}

type UserHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
