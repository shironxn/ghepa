package domain

import "gorm.io/gorm"

type Participant struct {
	gorm.Model
	UserID  uint  `json:"user_id" form:"user_id"`
	EventID uint  `json:"event_id" form:"event_id"`
	User    User  `json:"user" form:"user" gorm:"foreignKey:UserID;reference:ID"`
	Event   Event `json:"event" form:"event" gorm:"foreignKey:EventID;reference:ID"`
}
