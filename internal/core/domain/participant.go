package domain

import "gorm.io/gorm"

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
