package wordRepository

import (
	"gorm.io/gorm"
	"lingcard-go/models/word"
)

type WordRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *WordRepository {
	return &WordRepository{
		db: db,
	}
}

func (r *WordRepository) Find(id int) (word.Word, error) {
	var item word.Word
	err := r.db.Raw("SELECT * FROM words WHERE id = ?", id).Scan(&item).Error
	return item, err
}
