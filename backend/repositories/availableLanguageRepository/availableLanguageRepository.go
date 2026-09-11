package availableLanguageRepository

import (
	"gorm.io/gorm"
	"lingcard-go/models/language"
)

type AvailableLanguageRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AvailableLanguageRepository {
	return &AvailableLanguageRepository{
		db: db,
	}
}

func (r *AvailableLanguageRepository) FindAll() ([]language.AvailableLanguage, error) {
	var items []language.AvailableLanguage
	err := r.db.Raw("SELECT * FROM available_languages").Scan(&items).Error
	return items, err
}

func (r *AvailableLanguageRepository) FindByBaseLanguageId(id int) ([]language.AvailableLanguage, error) {
	var items []language.AvailableLanguage
	err := r.db.Raw("SELECT * FROM available_languages WHERE base_language_id = ?", id).Scan(&items).Error
	return items, err
}
