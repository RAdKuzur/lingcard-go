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

func (r *VoiceRepository) GetCountVoices(voteId int) int {
	var voices int
	r.db.Raw("SELECT COUNT(*) FROM voices"+" "+
		"JOIN vote_options ON vote_options.id = voices.vote_option_id"+" "+
		"JOIN votes ON votes.id = vote_options.vote_id"+" "+
		" WHERE votes.id = ?", voteId).Scan(&voices)
	return voices
}

func (r *VoiceRepository) FindUserVoice(userId int, voteOptionIds []int) int {
	var voice int
	r.db.Raw("SELECT vote_option_id FROM voices WHERE user_id = ? AND vote_option_id IN (?)", userId, voteOptionIds).Scan(&voice)
	return voice
}
