package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" validate:"required,max=30" gorm:"uniqueIndex"`
	Email    string `json:"email" form:"email" validate:"required,email" gorm:"uniqueIndex"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=100"`
}

type UserAuth struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}
