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

func (r *VoteRepository) All() []voice.Vote {
	var votes []voice.Vote
	r.db.Raw("SELECT * FROM votes").Scan(&votes)
	return votes
}

func (r *VoteRepository) Find(id int) voice.Vote {
	var Vote voice.Vote
	r.db.Raw("SELECT * FROM votes WHERE id = ?", id).Scan(&Vote)
	return Vote
}
