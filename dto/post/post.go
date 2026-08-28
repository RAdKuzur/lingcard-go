package post

import "time"

type SimplePostDTO struct {
	ID            int    `json:"id"`
	Content       string `json:"content"`
	Date          string `json:"date"`
	Title         string `json:"time"`
	Code          string `json:"code"`
	Username      string `json:"username"`
	Address       string `json:"address"`
	Status        string `json:"status"`
	ViewsCount    int    `json:"views_count"`
	LikesCount    int    `json:"likes_count"`
	DislikesCount int    `json:"dislikes_count"`
}

type PostDTO struct {
	ID            int          `json:"id"`
	Content       string       `json:"content"`
	Date          string       `json:"date"`
	Title         string       `json:"time"`
	Code          string       `json:"code"`
	Username      string       `json:"username"`
	Address       string       `json:"address"`
	Status        string       `json:"status"`
	ViewsCount    int          `json:"views_count"`
	LikesCount    int          `json:"likes_count"`
	DislikesCount int          `json:"dislikes_count"`
	IsLiked       bool         `json:"is_liked"`
	IsDisliked    bool         `json:"is_disliked"`
	Comments      []CommentDTO `json:"comments"`
}

type CommentDTO struct {
	ID           int    `json:"id"`
	Text         string `json:"text"`
	Username     string `json:"username"`
	Time         string `json:"time"`
	IsFixed      bool   `json:"is_fixed"`
	LanguageCode string `json:"language_code"`
}

type CreateCommentDTO struct {
	PostID  int    `json:"post_id"`
	UserID  int    `json:"user_id"`
	Text    string `json:"text"`
	Time    string `json:"time"`
	IsFixed bool   `json:"is_fixed"`
}

type CreatePostDTO struct {
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	LanguageID    int        `json:"language_id"`
	Address       string     `json:"address"`
	Status        int        `json:"status"`
	Date          *time.Time `json:"date"`
	ViewsCount    int        `json:"views_count"`
	LikesCount    int        `json:"likes_count"`
	DislikesCount int        `json:"dislikes_count"`
	UserID        int        `json:"user_id"`
}
