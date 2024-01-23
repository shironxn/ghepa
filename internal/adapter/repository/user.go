package repository

import (
	"errors"
	"event-planning-app/internal/core/domain"
	"event-planning-app/internal/core/port"

	"github.com/charmbracelet/log"
	"github.com/go-sql-driver/mysql"
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

func (u *UserRepository) Create(entity domain.User) (*domain.User, error) {
	err := u.db.Create(&entity).Error
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return nil, errors.New(mysqlErr.Message)
	}
	return &entity, err
}

func (u *UserRepository) GetAll() ([]domain.User, error) {
	var entity []domain.User
	err := u.db.Find(&entity).Error
	return entity, err
}

func (u *UserRepository) GetByID(id uint) (*domain.User, error) {
	var entity domain.User
	err := u.db.First(&entity, id).Error
	return &entity, err
}

func (u *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var entity domain.User
	err := u.db.Where("email = ?", email).First(&entity).Error
	return &entity, err
}

func (u *UserRepository) Update(entity *domain.User, entityUpdate domain.User) (*domain.User, error) {
	err := u.db.Model(&entity).Updates(entityUpdate).Error
	log.Info(entity)
	log.Info("TES")
	log.Info(entityUpdate)
	return entity, err
}

func (u *UserRepository) Delete(entity *domain.User) error {
	err := u.db.Delete(&entity).Error
	return err
}
