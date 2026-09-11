package voteOptionRepository

import (
	"gorm.io/gorm"
	"lingcard-go/models/voice"
)

type VoteOptionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *VoteOptionRepository {
	return &VoteOptionRepository{
		db: db,
	}
}
func (r *VoteOptionRepository) Find(id int) (voice.VoteOption, error) {
	var voteOption voice.VoteOption
	err := r.db.Raw("SELECT * FROM vote_options WHERE id = ?", id).Scan(&voteOption).Error
	return voteOption, err
}

func (r *VoteOptionRepository) FindByVoteId(id int) ([]voice.VoteOption, error) {
	var voteOptions []voice.VoteOption
	err := r.db.Raw("SELECT * FROM vote_options WHERE vote_id = ?", id).Scan(&voteOptions).Error
	return voteOptions, err
}
