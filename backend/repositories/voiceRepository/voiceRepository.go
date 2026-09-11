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

func (r *VoiceRepository) GetCountVoices(voteId int) (int, error) {
	var voices int
	err := r.db.Raw("SELECT COUNT(*) FROM voices"+" "+
		"JOIN vote_options ON vote_options.id = voices.vote_option_id"+" "+
		"JOIN votes ON votes.id = vote_options.vote_id"+" "+
		" WHERE votes.id = ?", voteId).Scan(&voices).Error
	return voices, err
}

func (r *VoiceRepository) FindUserVoice(userId int, voteOptionIds []int) (voice.Voice, error) {
	var item voice.Voice
	err := r.db.Raw("SELECT * FROM voices WHERE user_id = ? AND vote_option_id IN (?) LIMIT 1", userId, voteOptionIds).Scan(&item).Error
	return item, err
}

func (r *VoiceRepository) DeleteUserVoices(userId int, voteOptionIds []int) error {
	err := r.db.Exec("DELETE FROM voices WHERE user_id = ? AND vote_option_id IN (?)", userId, voteOptionIds).Error
	return err
}

func (r *VoiceRepository) CountVoices(voteOptionId int) (int, error) {
	var count int
	err := r.db.Raw("SELECT count(*) FROM voices WHERE vote_option_id = ?", voteOptionId).Scan(&count).Error
	return count, err
}

func (r *VoiceRepository) DeleteVoice(userID int, voteOptionID int) error {
	err := r.db.Exec("DELETE FROM voices WHERE user_id = ? AND vote_option_id = ?", userID, voteOptionID).Error
	return err
}

func (r *VoiceRepository) Insert(dto voiceDTO.VoiceDTO) error {
	err := r.db.Exec("INSERT INTO voices (vote_option_id, user_id, time) VALUES (?, ?, ?)", dto.VoteOptionID, dto.UserID, dto.Time).Error
	return err
}
