package wordTranslationRepository

import (
	"gorm.io/gorm"
	"lingcard-go/models/words"
)

type WordTranslationRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *WordTranslationRepository {
	return &WordTranslationRepository{
		db: db,
	}
}

func (r *WordTranslationRepository) GetPaginateByTargetLanguageIdAndBaseLanguageId(baseLangId, targetLangId, page, limit int, search string) []words.WordTranslationDTO {
	var items []words.WordTranslationDTO
	r.db.Raw("SELECT * FROM word_translations JOIN words ON word_translations.word_id = words.id WHERE word_translations.target_language_id = ? words.language_id = ? WHERE words.text LIKE %"+search+"%", baseLangId, targetLangId).Scan(&items)
	return items

}

func (r *WordTranslationRepository) CountByTargetLanguageIdAndBaseLanguageId(baseLangId, targetLangId int, search string) int {
	var count int
	r.db.Raw("SELECT COUNT(*) FROM word_translations JOIN words ON word_translations.word_id = words.id WHERE word_translations.target_language_id = ? words.language_id = ? WHERE words.text LIKE %"+search+"%", baseLangId, targetLangId).Scan(&count)
	return count
}
