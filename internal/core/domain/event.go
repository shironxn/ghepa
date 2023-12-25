package domain

import (
	"event-planning-app/internal/response"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title       string        `json:"title" form:"title" validate:"required"`
	Description string        `json:"description" form:"description" validate:"required"`
	EndDate     string        `json:"end_date" form:"end_date" validate:"required,datetime=2006-01-02"`
	UserID      uint          `json:"user_id" form:"user_id"`
	User        response.User `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	Comments    []Comment     `json:"comment" form:"comment" gorm:"many2many:comment_event"`
	Participant []Participant `json:"participant" form:"participant" gorm:"many2many:participant_event"`
}
