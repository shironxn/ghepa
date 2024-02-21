package domain

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	EndDate      string        `json:"end_date"`
	UserID       uint          `json:"user_id"`
	User         UserResponse  `json:"user" gorm:"foreignKey:UserID;reference:ID"`
	Comments     []Comment     `json:"comments" gorm:"many2many:comments_event"`
	Participants []Participant `json:"participants" gorm:"many2many:participants_event"`
}

type EventRequest struct {
	User        UserResponse `json:"user"`
	Title       string       `json:"title" form:"title" validate:"required"`
	Description string       `json:"description" form:"description" validate:"required"`
	EndDate     string       `json:"end_date" form:"end_date" validate:"required,datetime=2006-01-02"`
}

type EventResponse struct {
	ID           uint              `json:"id"`
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	Owner        UserResponse      `json:"owner"`
	EndDate      string            `json:"end_date"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	Comments     []CommentList     `json:"comments"`
	Participants []ParticipantList `json:"participants"`
}
