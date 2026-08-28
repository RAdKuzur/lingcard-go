package word

type WordTranslation struct {
	ID               int    `gorm:"column:id"`
	Translation      string `gorm:"column:translation"`
	WordID           int    `gorm:"column:word_id"`
	TargetLanguageID int    `gorm:"column:target_language_id"`
}

type Word struct {
	ID            int    `gorm:"column:id"`
	Text          string `gorm:"column:text"`
	Transcription string `gorm:"column:transcription"`
	LanguageID    int    `gorm:"column:language_id"`
	Level         int    `gorm:"column:level"`
}
