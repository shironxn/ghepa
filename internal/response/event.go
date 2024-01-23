package response

import "time"

type Event struct {
	ID           uint              `json:"id"`
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	Owner        User              `json:"owner"`
	EndDate      string            `json:"end_date"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	Comments     []CommentList     `json:"comments"`
	Participants []ParticipantList `json:"participants"`
}
