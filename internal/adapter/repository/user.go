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

func (u *UserRepository) Create(req domain.User) (*domain.User, error) {
	err := u.db.Create(&req).Error
	return &req, err
}

func (u *UserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := u.db.Find(&users).Error
	return users, err
}

func (u *UserRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	err := u.db.First(&user, id).Error
	return &user, err
}

func (u *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := u.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u *UserRepository) Update(user *domain.User, req domain.User) (*domain.User, error) {
	err := u.db.Model(user).Updates(req).Error
	return user, err
}

func (u *UserRepository) Delete(user *domain.User) error {
	err := u.db.Delete(&user).Error
	return err
}
