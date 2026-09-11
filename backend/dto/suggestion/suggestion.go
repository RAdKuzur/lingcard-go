package suggestion

import "time"

type SuggestionDTO struct {
	UserID  int        `json:"user_id"`
	Message string     `json:"message"`
	Date    *time.Time `json:"date"`
	Status  bool       `json:"status"`
}
