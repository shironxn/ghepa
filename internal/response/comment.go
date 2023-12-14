package response

import "time"

type CommentResponse struct {
	UserName  string    `json:"user_name"`
	EventName string    `json:"event_name"`
	Comment   string    `json:"comment"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}
