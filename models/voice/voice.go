package voice

import "time"

type Vote struct {
	ID       int    `gorm:"column:id"`
	Title    string `gorm:"column:title"`
	Content  string `gorm:"column:content"`
	IsActive bool   `gorm:"column:is_active"`
}

type VoteOption struct {
	ID      int    `gorm:"column:id"`
	VoteID  int    `gorm:"column:vote_id"`
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

type Voice struct {
	ID           int        `gorm:"column:id"`
	UserID       int        `gorm:"column:user_id"`
	VoteOptionID int        `gorm:"column:vote_option_id"`
	Time         *time.Time `gorm:"column:time"`
}
