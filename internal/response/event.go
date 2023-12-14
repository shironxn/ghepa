package response

type EventResponse struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	User        string            `json:"user"`
	Comments    []CommentResponse `json:"comment"`
	Participant []CommentResponse `json:"participant"`
}
