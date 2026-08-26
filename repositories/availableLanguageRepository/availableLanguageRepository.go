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

func (r *AvailableLanguageRepository) All() []language.AvailableLanguage {
	var languages []language.AvailableLanguage
	_ = r.db.Find(&languages).Error
	return languages
}

func (r *AvailableLanguageRepository) FindByBaseLanguageId(id int) []language.AvailableLanguage {
	var l []language.AvailableLanguage
	_ = r.db.Preload("TargetLanguage").
		Where("base_language_id = ?", id).Find(&l)
	return l
}
