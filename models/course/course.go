package course

type Course struct {
	ID                int    `gorm:"column:id"`
	WordTranslationID int    `gorm:"column:word_translation_id"`
	UserID            int    `gorm:"column:user_id"`
	Repeat            int    `gorm:"column:repeat"`
	Status            string `gorm:"column:status"`
	LastTimeRepeated  int    `gorm:"column:last_time_repeated"`
}
