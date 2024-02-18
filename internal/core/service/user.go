package service

import (
	"errors"
	"event-planning-app/internal/core/domain"
	"event-planning-app/internal/core/port"
	"event-planning-app/internal/util"

	"github.com/charmbracelet/log"
)

type UserService struct {
	repository port.UserRepository
	util       util.Bcrypt
}

func NewUserService(repository port.UserRepository) port.UserService {
	return &UserService{
		repository: repository,
	}
}

func (u *UserService) Create(entity domain.RegisterRequest) (*domain.User, error) {
	hashedPassword, err := u.util.HashPassword(entity.Password)
	if err != nil {
		return nil, err
	}

	entity.Password = string(hashedPassword)

	data, err := u.repository.Create(entity)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) Login(entity domain.LoginRequest) (*domain.User, error) {
	data, err := u.repository.GetByEmail(entity.Email)
	if err != nil {
		return nil, err
	}

	if err := u.util.ComparePassword(entity.Password, []byte(data.Password)); err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) GetAll() ([]domain.User, error) {
	data, err := u.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) GetByID(id uint) (*domain.User, error) {
	data, err := u.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) Update(entity domain.User, id uint) (*domain.User, error) {
	user, err := u.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	log.Info(entity.ID)

	if user.ID != entity.ID {
		return nil, errors.New("user does not have permission to perform this action")
	}

	hashedPassword, err := u.util.HashPassword(entity.Password)
	if err != nil {
		return nil, err
	}

	entity.Password = string(hashedPassword)

	data, err := u.repository.Update(user, entity)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) Delete(entity domain.User) error {
	user, err := u.repository.GetByID(entity.ID)
	if err != nil {
		return err
	}

	if user.Name != entity.Name {
		return errors.New("user does not have permission to perform this action")
	}

	err = u.repository.Delete(user)
	if err != nil {
		return err
	}

	return nil

}
