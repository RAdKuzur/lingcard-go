package suggestionRepository

import (
	"gorm.io/gorm"
	"lingcard-go/dto/suggestion"
	"time"
)

type SuggestionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *SuggestionRepository {
	return &SuggestionRepository{
		db: db,
	}
}

func (r *SuggestionRepository) Insert(dto suggestion.SuggestionDTO) error {
	err := r.db.Exec("INSERT INTO suggestions (message, date, user_id, status) VALUES (?, ?, ?, ?)", dto.Message, time.Now(), dto.UserID, false).Error
	return err
}
