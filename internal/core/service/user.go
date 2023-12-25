package service

import (
	"errors"
	"event-planning-app/internal/core/domain"
	"event-planning-app/internal/core/port"
	"event-planning-app/internal/util"
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

func (u *UserService) Create(req domain.User) (*domain.User, error) {
	if _, err := u.repository.GetByEmail(req.Email); err == nil {
		return nil, errors.New("email has already been used")
	}

	hashedPassword, err := u.util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashedPassword)

	data, err := u.repository.Create(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) Login(req domain.UserAuth) (*domain.User, error) {
	data, err := u.repository.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if err := u.util.ComparePassword(req.Password, []byte(data.Password)); err != nil {
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

func (u *UserService) Update(id uint, req domain.User) (*domain.User, error) {
	hashedPassword, err := u.util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashedPassword)

	user, err := u.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	data, err := u.repository.Update(user, req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) Delete(id uint) error {
	user, err := u.repository.GetByID(id)
	if err != nil {
		return err
	}

	err = u.repository.Delete(user)
	if err != nil {
		return err
	}

	return nil

}
