package voiceRepository

import (
	"gorm.io/gorm"
	voiceDTO "lingcard-go/dto/voice"
	"lingcard-go/models/voice"
)

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

func (r *VoiceRepository) FindUserVoice(userId int, voteOptionIds []int) voice.Voice {
	var item voice.Voice
	r.db.Raw("SELECT * FROM voices WHERE user_id = ? AND vote_option_id IN (?) LIMIT 1", userId, voteOptionIds).Scan(&item)
	return item
}

func (r *VoiceRepository) DeleteUserVoices(userId int, voteOptionIds []int) {
	r.db.Exec("DELETE FROM voices WHERE user_id = ? AND vote_option_id IN (?)", userId, voteOptionIds)
}

func (r *VoiceRepository) CountVoices(voteOptionId int) int {
	var count int
	r.db.Raw("SELECT count(*) FROM voices WHERE vote_option_id = ?", voteOptionId).Scan(&count)
	return count
}

func (r *VoiceRepository) DeleteVoice(userID int, voteOptionID int) {
	r.db.Exec("DELETE FROM voices WHERE user_id = ? AND vote_option_id = ?", userID, voteOptionID)
}

func (r *VoiceRepository) Insert(dto voiceDTO.VoiceDTO) {
	r.db.Exec("INSERT INTO voices (vote_option_id, user_id, time)", dto.VoteOptionID, dto.UserID, dto.Time)
}
