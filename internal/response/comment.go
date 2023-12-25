package response

import "time"

type Comment struct {
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
