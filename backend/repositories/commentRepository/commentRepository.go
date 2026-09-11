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

func (r *CommentRepository) Find(postID int) (post.Comment, error) {
	var comment post.Comment
	err := r.db.Raw("SELECT * FROM comments WHERE id = ?", postID).Scan(&comment).Error
	return comment, err
}

func (r *CommentRepository) GetAllComments(postID int) ([]post.Comment, error) {
	var comments []post.Comment
	err := r.db.Raw("SELECT * FROM comments WHERE post_id = ? ORDER BY is_fixed, time desc", postID).Scan(&comments).Error
	return comments, err
}

func (r *CommentRepository) Insert(commentDTO postDTO.CreateCommentDTO) error {
	err := r.db.Exec("INSERT INTO comments (post_id, user_id, text, time, is_fixed) VALUES (?, ?, ?, ?, ?)", commentDTO.PostID, commentDTO.UserID, commentDTO.Text, time.Now(), false).Error
	return err
}

func (r *CommentRepository) Delete(id int) error {
	err := r.db.Exec("DELETE FROM comments WHERE id = ?", id).Error
	return err
}
