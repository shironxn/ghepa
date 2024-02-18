package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"uniqueIndex"`
	Email    string `json:"email" form:"email" gorm:"uniqueIndex"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name" form:"name" validate:"required,min=4,max=30"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=100"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}
