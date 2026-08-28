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

func (r *WordRepository) Find(id int) word.Word {
	var item word.Word
	r.db.Raw("SELECT * FROM words WHERE id = ?", id).Scan(&item)
	return item
}
