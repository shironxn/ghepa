package domain

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	EndDate      string        `json:"end_date"`
	UserID       uint          `json:"user_id"`
	User         UserResponse  `json:"user" gorm:"foreignKey:UserID;reference:ID"`
	Comments     []Comment     `json:"comments" gorm:"many2many:comment_events"`
	Participants []Participant `json:"participants" gorm:"many2many:participant_events"`
}

type EventRequest struct {
	User        UserResponse `json:"user"`
	Name        string       `json:"Name" form:"Name" validate:"required"`
	Description string       `json:"description" form:"description" validate:"required"`
	EndDate     string       `json:"end_date" form:"end_date" validate:"required,datetime=2006-01-02"`
}

type EventResponse struct {
	ID           uint              `json:"id"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	UserID       uint              `json:"user_id"`
	UserName     string            `json:"user_name"`
	EndDate      string            `json:"end_date"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	Comments     []CommentList     `json:"comments,omitempty"`
	Participants []ParticipantList `json:"participants,omitempty"`
}
