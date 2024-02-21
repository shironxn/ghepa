package domain

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	EventID uint   `json:"event_id"`
	Comment string `json:"comment"`
	User    User   `json:"user" gorm:"foreignKey:UserID;reference:ID"`
	Event   Event  `json:"event" gorm:"foreignKey:EventID;reference:ID"`
}

type CommentRequest struct {
	UserID  uint   `json:"user_id" form:"user_id" validate:"required"`
	EventID uint   `json:"event_id" form:"event_id" validate:"required"`
	Comment string `json:"comment" form:"comment" validate:"required"`
}

type CommentResponse struct {
	UserName  string    `json:"user_name"`
	EventName string    `json:"event_name"`
	Comment   string    `json:"comment"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type CommentList struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}
