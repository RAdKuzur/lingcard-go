package language

type Language struct {
	ID       int    `gorm:"column:id;primaryKey"`
	Name     string `gorm:"column:name"`
	Code     string `gorm:"column:code"`
	IsActive bool   `gorm:"column:is_active"`
}

type AvailableLanguage struct {
	ID               int `gorm:"column:id;primaryKey"`
	BaseLanguageID   int `gorm:"column:base_language_id"`
	TargetLanguageID int `gorm:"column:target_language_id"`
}
