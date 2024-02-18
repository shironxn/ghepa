package domain

import (
	"event-planning-app/internal/response"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint           `json:"user_id" form:"user_id" validate:"required"`
	EventID uint           `json:"event_id" form:"event_id" validate:"required"`
	Comment string         `json:"comment" form:"comment" validate:"required"`
	User    response.User  `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	Event   response.Event `json:"event" form:"event" gorm:"foreignKey:EventID;reference:ID"`
}
