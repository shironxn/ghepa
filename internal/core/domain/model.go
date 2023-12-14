package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type Event struct {
	gorm.Model
	Name        string        `json:"name" form:"name" validate:"required"`
	Description string        `json:"description" form:"description" validate:"required"`
	UserID      uint          `json:"user_id" form:"user_id" validate:"required"`
	User        User          `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	Comments    []Comment     `json:"comment" form:"comment" gorm:"many2many:comment_event"`
	Participant []Participant `json:"participant" form:"participant" gorm:"many2many:participant_event"`
}

type Comment struct {
	gorm.Model
	UserID  uint   `json:"user_id" form:"user_id" validate:"required"`
	EventID uint   `json:"event_id" form:"event_id" validate:"required"`
	Comment string `json:"comment" form:"comment" validate:"required"`
	User    User   `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	Event   Event  `json:"event" form:"event" gorm:"foreignKey:EventID;reference:ID"`
}

type Participant struct {
	gorm.Model
	UserID  uint  `json:"user_id" form:"user_id"`
	EventID uint  `json:"event_id" form:"event_id"`
	User    User  `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	Event   Event `json:"event" form:"event" gorm:"foreignKey:EventID;reference:ID"`
}
