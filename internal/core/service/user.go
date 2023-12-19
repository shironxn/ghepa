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

func (us *UserService) Create(req domain.User) (*domain.User, error) {
	if checkEmail, _ := us.repository.GetByEmail(req.Email); checkEmail != nil {
		return nil, errors.New("email has already been used")
	}

	hashedPassword, err := us.util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashedPassword)

	data, err := us.repository.Create(req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (us *UserService) Login(req domain.UserAuth) (*domain.User, error) {
	data, err := us.repository.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if err := us.util.ComparePassword(req.Password, []byte(data.Password)); err != nil {
		return nil, err
	}

	return data, nil
}

func (us *UserService) GetAll() ([]domain.User, error) {
	data, err := us.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (us *UserService) GetByID(id uint) (*domain.User, error) {
	data, err := us.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (us *UserService) Update(id uint, req domain.User) (*domain.User, error) {
	hashedPassword, err := us.util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashedPassword)

	user, err := us.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	data, err := us.repository.Update(*user, req)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (us *UserService) Delete(id uint) error {
	err := us.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil

}
