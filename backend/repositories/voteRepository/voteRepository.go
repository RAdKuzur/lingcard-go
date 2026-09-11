package voteRepository

import (
	"gorm.io/gorm"
	"lingcard-go/models/voice"
)

type VoteRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *VoteRepository {
	return &VoteRepository{
		db: db,
	}
}

func (r *VoteRepository) All() ([]voice.Vote, error) {
	var votes []voice.Vote
	err := r.db.Raw("SELECT * FROM votes").Scan(&votes).Error
	return votes, err
}

func (r *VoteRepository) Find(id int) (voice.Vote, error) {
	var item voice.Vote
	err := r.db.Raw("SELECT * FROM votes WHERE id = ?", id).Scan(&item).Error
	return item, err
}
