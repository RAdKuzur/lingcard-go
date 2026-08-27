package voiceRepository

import "gorm.io/gorm"

type VoiceRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *VoiceRepository {
	return &VoiceRepository{
		db: db,
	}
}
