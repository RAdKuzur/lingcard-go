package languageRepository

import (
	"gorm.io/gorm"
	"lingcard-go/models/language"
)

type LanguageRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *LanguageRepository {
	return &LanguageRepository{
		db: db,
	}
}

func (r *LanguageRepository) All() ([]language.Language, error) {
	var languages []language.Language
	err := r.db.Raw("SELECT * FROM languages").Scan(&languages).Error
	return languages, err
}

func (r *LanguageRepository) AllActive() ([]language.Language, error) {
	var languages []language.Language
	err := r.db.Raw("SELECT * FROM languages WHERE is_active = ?", true).Scan(&languages).Error
	return languages, err
}

func (r *LanguageRepository) FindByCode(code string) (language.Language, error) {
	var lang language.Language
	err := r.db.Raw("SELECT * FROM languages WHERE code = ?", code).Scan(&lang).Error
	return lang, err
}

func (r *LanguageRepository) Find(id int) (language.Language, error) {
	var lang language.Language
	err := r.db.Raw("SELECT * FROM languages WHERE id = ?", id).Scan(&lang).Error
	return lang, err
}
