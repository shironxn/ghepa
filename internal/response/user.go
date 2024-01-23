package response

import "time"

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserDetails struct {
	Token   string `json:"token,omitempty"`
	Expired string `json:"expired_at,omitempty"`
}
