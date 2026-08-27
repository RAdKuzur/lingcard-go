package voteOptionRepository

import (
	"gorm.io/gorm"
	"lingcard-go/dto/voice"
)

type VoteOptionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *VoteOptionRepository {
	return &VoteOptionRepository{
		db: db,
	}
}

func (r *VoteOptionRepository) FindByVoteId(id int) []voice.VoteOptionDTO {
	var voteOptions []voice.VoteOptionDTO
	r.db.Raw("SELECT * FROM vote_options WHERE vote_id = ?", id).Find(&voteOptions)
	return voteOptions
}
