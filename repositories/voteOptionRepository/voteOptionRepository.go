package voteOptionRepository

import "gorm.io/gorm"

type VoteOptionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *VoteOptionRepository {
	return &VoteOptionRepository{
		db: db,
	}
}
