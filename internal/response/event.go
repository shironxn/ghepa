package response

import "time"

type Event struct {
	ID          uint              `json:"event_id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Owner       User              `json:"owner"`
	Comments    []CommentList     `json:"comment"`
	Participant []ParticipantList `json:"participant"`
}

type EventDetails struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	EndDate   string    `json:"end_date"`
}

type EventInfo struct {
	Event        Event        `json:"event"`
	EventDetails EventDetails `json:"details"`
}
