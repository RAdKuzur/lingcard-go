package word

type WordTranslation struct {
	ID            int    `gorm:"column:id"`
	Text          string `gorm:"column:text"`
	Translation   string `gorm:"column:translation"`
	Transcription string `gorm:"column:transcription"`
	Level         int    `gorm:"column:level"`
}
