package port

import (
	"event-planning-app/internal/core/domain"
	"net/http"
)

type UserRepository interface {
	Create(req domain.User) (*domain.User, error)
	GetAll() ([]domain.User, error)
	GetByID(id uint) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Update(user *domain.User, req domain.User) (*domain.User, error)
	Delete(user *domain.User) error
}

type UserService interface {
	Create(req domain.User) (*domain.User, error)
	Login(req domain.UserAuth) (*domain.User, error)
	GetAll() ([]domain.User, error)
	GetByID(id uint) (*domain.User, error)
	Update(id uint, req domain.User) (*domain.User, error)
	Delete(id uint) error
}

type UserHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
