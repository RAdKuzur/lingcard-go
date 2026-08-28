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

func (r *LanguageRepository) All() []language.Language {
	var languages []language.Language
	_ = r.db.Find(&languages).Error
	return languages
}

func (r *LanguageRepository) AllActive() []language.Language {
	var languages []language.Language
	_ = r.db.Preload("BaseLanguages.TargetLanguage").
		Where("is_active = ?", true).
		Find(&languages).Error
	return languages
}

func (r *LanguageRepository) FindByCode(code string) language.Language {
	var Language language.Language
	_ = r.db.First(&Language, "code = ?", code).Error
	return Language
}

func (r *LanguageRepository) Find(id int) language.Language {
	var Language language.Language
	_ = r.db.First(&Language, "id = ?", id).Error
	return Language
}
