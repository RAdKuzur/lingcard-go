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
func (r *WordTranslationRepository) Find(id int) (word.WordTranslation, error) {
	var item word.WordTranslation
	err := r.db.Raw("SELECT * FROM word_translations WHERE id = ?", id).Scan(&item).Error
	return item, err
}

func (r *WordTranslationRepository) GetPaginateByTargetLanguageIdAndBaseLanguageId(baseLangId, targetLangId, page, limit int, search string) ([]word.WordTranslation, error) {
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
	err := r.db.Raw(query, args...).Scan(&items).Error
	return items, err
}

func (r *WordTranslationRepository) CountByTargetLanguageIdAndBaseLanguageId(baseLangId, targetLangId int, search string) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM word_translations JOIN words ON word_translations.word_id = words.id WHERE word_translations.target_language_id = ? AND words.language_id = ?"
	if search != "" {
		query = query + " AND word.text LIKE %" + search + "%"
	}

	err := r.db.Raw(query, baseLangId, targetLangId).Scan(&count).Error
	return count, err
}

func (r *WordTranslationRepository) CountNewWords(baseLangId, targetLangId int, coursesId []int) (int, error) {
	var count int

	query := r.db.Table("word_translations").
		Select("COUNT(*)").
		Joins("JOIN words ON word_translations.word_id = words.id").
		Where("word_translations.target_language_id = ?", baseLangId).
		Where("words.language_id = ?", targetLangId)

	if len(coursesId) > 0 {
		query = query.Where("words.id NOT IN (?)", coursesId)
	}

	err := query.Scan(&count).Error
	return count, err
}

func (r *WordTranslationRepository) GetSearchNewWords(baseLangId, targetLangId, page, limit int, search string, exceptID []int) ([]word.WordTranslation, error) {

	var words []word.WordTranslation

	query := r.db.Table("word_translations").
		Select("word_translations.*").
		Joins("JOIN words ON word_translations.word_id = words.id").
		Where("word_translations.target_language_id = ?", baseLangId).
		Where("words.language_id = ?", targetLangId)

	if search != "" {
		query = query.Where("words.text LIKE ?", "%"+search+"%")
	}

	if len(exceptID) > 0 {
		query = query.Where("words.id NOT IN (?)", exceptID)
	}
	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Find(&words).Error
	return words, err

}

func (r *WordTranslationRepository) CountSearchNewWords(baseLangId, targetLangId int, search string, exceptID []int) (int, error) {

	var count int

	query := r.db.Table("word_translations").
		Select("COUNT(*)").
		Joins("JOIN words ON word_translations.word_id = words.id").
		Where("word_translations.target_language_id = ?", baseLangId).
		Where("words.language_id = ?", targetLangId)

	if search != "" {
		query = query.Where("words.text LIKE ?", "%"+search+"%")
	}

	if len(exceptID) > 0 {
		query = query.Where("words.id NOT IN (?)", exceptID)
	}
	err := query.Scan(&count).Error
	return count, err
}
