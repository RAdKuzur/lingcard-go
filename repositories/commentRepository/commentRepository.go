package commentRepository

import (
	"gorm.io/gorm"
	postDTO "lingcard-go/dto/post"
	"lingcard-go/models/post"
	"time"
)

type CommentRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) Find(postID int) post.Comment {
	var comment post.Comment
	r.db.Raw("SELECT * FROM comments WHERE id = ?", postID).Scan(&comment)
	return comment
}

func (r *CommentRepository) GetAllComments(postID int) []post.Comment {
	var comments []post.Comment
	r.db.Raw("SELECT * FROM comments WHERE post_id = ? ORDER BY is_fixed, time desc", postID).Scan(&comments)
	return comments
}

func (r *CommentRepository) Insert(commentDTO postDTO.CreateCommentDTO) {
	r.db.Exec("INSERT INTO comments (post_id, user_id, text, time, is_fixed) VALUES (?, ?, ?, ?, ?)", commentDTO.PostID, commentDTO.UserID, commentDTO.Text, time.Now(), false)
}

func (r *CommentRepository) Delete(id int) {
	r.db.Exec("DELETE FROM comments WHERE id = ?", id)
}
