package voteRepository

import "gorm.io/gorm"

type VoteRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *VoteRepository {
	return &VoteRepository{
		db: db,
	}
}
