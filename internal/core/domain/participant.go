package domain

import (
	"time"

	"gorm.io/gorm"
)

type Participant struct {
	gorm.Model
	UserID  uint  `json:"user_id"`
	EventID uint  `json:"event_id"`
	User    User  `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	Event   Event `json:"event" form:"event" gorm:"foreignKey:EventID;reference:ID"`
}

type ParticipantRequest struct {
	UserID  uint `json:"user_id" validate:"required"`
	EventID uint `json:"event_id" validate:"required"`
}

type ParticipantList struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ParticipantResponse struct {
	ID        uint      `json:"id"`
	EventID   uint      `json:"event_id"`
	UserID    uint      `json:"user_id"`
	EventName string    `json:"event_name"`
	UserName  string    `json:"user_name"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}
