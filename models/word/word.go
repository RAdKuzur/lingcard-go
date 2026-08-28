package word

type WordTranslation struct {
	ID            int    `gorm:"column:id"`
	Text          string `gorm:"column:text"`
	Translation   string `gorm:"column:translation"`
	Transcription string `gorm:"column:transcription"`
	Level         int    `gorm:"column:level"`
}

type Word struct {
	ID            int    `gorm:"column:id"`
	Text          string `gorm:"column:text"`
	Transcription string `gorm:"column:transcription"`
	LanguageID    int    `gorm:"column:language_id"`
	Level         int    `gorm:"column:level"`
}
