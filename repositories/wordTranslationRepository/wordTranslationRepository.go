package wordTranslationRepository

import (
	"gorm.io/gorm"
	"lingcard-go/models/word"
)

type WordTranslationRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *WordTranslationRepository {
	return &WordTranslationRepository{
		db: db,
	}
}

func (r *WordTranslationRepository) GetPaginateByTargetLanguageIdAndBaseLanguageId(baseLangId, targetLangId, page, limit int, search string) []word.WordTranslation {
	var items []word.WordTranslation
	query := "SELECT word_translations.*, words.* FROM word_translations JOIN words ON word_translations.word_id = words.id WHERE word_translations.target_language_id = ? AND words.language_id = ?"
	args := []interface{}{baseLangId, targetLangId}
	if search != "" {
		query += " AND word.text LIKE ?"
		args = append(args, "%"+search+"%")
	}
	offset := (page - 1) * limit
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)
	r.db.Raw(query, args...).Scan(&items)
	return items
}

func (r *WordTranslationRepository) CountByTargetLanguageIdAndBaseLanguageId(baseLangId, targetLangId int, search string) int {
	var count int
	query := "SELECT COUNT(*) FROM word_translations JOIN words ON word_translations.word_id = words.id WHERE word_translations.target_language_id = ? AND words.language_id = ?"
	if search != "" {
		query = query + " AND word.text LIKE %" + search + "%"
	}

	r.db.Raw(query, baseLangId, targetLangId).Scan(&count)
	return count
}
