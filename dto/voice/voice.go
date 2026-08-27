package voice

type VoteOptionDTO struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Count   string `json:"count"`
}

type VoteDTO struct {
	ID          int             `json:"id"`
	Title       string          `json:"title"`
	Content     string          `json:"content"`
	VoteOptions []VoteOptionDTO `json:"vote_options"`
	TotalCount  int             `json:"total_count"`
	IsActive    bool            `json:"is_active"`
	Voted       *int            `json:"voted"`
}

type SimpleVoteDTO struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Voices  int    `json:"voices"`
}
