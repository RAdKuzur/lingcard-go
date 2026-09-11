package post

import "time"

type Post struct {
	ID            int        `gorm:"column:id"`
	Title         string     `gorm:"column:title"`
	Content       string     `gorm:"column:content"`
	LanguageID    int        `gorm:"column:language_id"`
	Date          *time.Time `gorm:"column:date"`
	UserID        int        `gorm:"column:user_id"`
	Address       string     `gorm:"column:address"`
	Status        int        `gorm:"column:status"`
	ViewsCount    int        `gorm:"column:views_count"`
	LikesCount    int        `gorm:"column:likes_count"`
	DislikesCount int        `gorm:"column:dislikes_count"`
}

type Comment struct {
	ID      int        `gorm:"column:id"`
	PostID  int        `gorm:"column:post_id"`
	UserID  int        `gorm:"column:user_id"`
	Text    string     `gorm:"column:text"`
	Time    *time.Time `gorm:"column:time"`
	IsFixed bool       `gorm:"column:is_fixed"`
}
