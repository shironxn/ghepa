package repository

import (
	"event-planning-app/internal/core/domain"
	"event-planning-app/internal/core/port"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(req domain.User) (*domain.User, error) {
	err := ur.db.Create(&req).Error
	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (ur *UserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User

	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User

	err := ur.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) Update(user domain.User, req domain.User) (*domain.User, error) {
	err := ur.db.Model(&user).Updates(req).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) Delete(id uint) error {
	var user domain.User

	err := ur.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}
