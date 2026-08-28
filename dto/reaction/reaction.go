package reaction

type ReactionDTO struct {
	UserID int `json:"user_id"`
	PostID int `json:"post_id"`
	Status int `json:"status"`
}
