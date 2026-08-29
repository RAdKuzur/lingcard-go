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
func (r *VoteOptionRepository) Find(id int) voice.VoteOption {
	var voteOption voice.VoteOption
	r.db.Raw("SELECT * FROM vote_options WHERE id = ?", id).Scan(&voteOption)
	return voteOption
}

func (r *VoteOptionRepository) FindByVoteId(id int) []voice.VoteOption {
	var voteOptions []voice.VoteOption
	r.db.Raw("SELECT * FROM vote_options WHERE vote_id = ?", id).Scan(&voteOptions)
	return voteOptions
}
