package language

type Language struct {
	Id              int                 `gorm:"column:id;primaryKey"`
	Name            string              `gorm:"column:name"`
	Code            string              `gorm:"column:code"`
	IsActive        bool                `gorm:"column:is_active"`
	BaseLanguages   []AvailableLanguage `gorm:"foreignKey:BaseLanguageId;references:Id"`
	TargetLanguages []AvailableLanguage `gorm:"foreignKey:TargetLanguageId;references:Id"`
}

type AvailableLanguage struct {
	Id               int      `gorm:"column:id;primaryKey"`
	BaseLanguageId   int      `gorm:"column:base_language_id"`
	TargetLanguageId int      `gorm:"column:target_language_id"`
	BaseLanguage     Language `gorm:"foreignKey:BaseLanguageId;references:Id"`
	TargetLanguage   Language `gorm:"foreignKey:TargetLanguageId;references:Id"`
}
