package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type UserAuth struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}
