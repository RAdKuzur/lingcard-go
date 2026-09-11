package postRepository

import (
	"gorm.io/gorm"
	postDict "lingcard-go/dictionaries/post"
	postDTO "lingcard-go/dto/post"
	"lingcard-go/models/post"
	"time"
)

type PostRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}
func (r *PostRepository) Find(id int) (post.Post, error) {
	var item post.Post
	err := r.db.Raw("SELECT * FROM posts WHERE id = ?", id).Scan(&item).Error
	return item, err
}

func (r *PostRepository) FindApprovedPostsByLanguageId(langId int) ([]post.Post, error) {
	var posts []post.Post
	err := r.db.Raw("SELECT * FROM posts WHERE language_id = ? AND status = ? ORDER BY date ASC", langId, postDict.APPROVED).Scan(&posts).Error
	return posts, err
}

func (r *PostRepository) IncrementViewsCount(postId int) error {
	err := r.db.Exec("UPDATE posts SET views_count = views_count + 1 WHERE id = ?", postId).Error
	return err
}

func (r *PostRepository) IncrementLikesCount(postId int) error {
	err := r.db.Exec("UPDATE posts SET likes_count = likes_count + 1 WHERE id = ?", postId).Error
	return err
}

func (r *PostRepository) IncrementDislikesCount(postId int) error {
	err := r.db.Exec("UPDATE posts SET dislikes_count = dislikes_count + 1 WHERE id = ?", postId).Error
	return err
}

func (r *PostRepository) DecrementLikesCount(postId int) error {
	err := r.db.Exec("UPDATE posts SET likes_count = likes_count - 1 WHERE id = ?", postId).Error
	return err
}

func (r *PostRepository) DecrementDislikesCount(postId int) error {
	err := r.db.Exec("UPDATE posts SET dislikes_count = dislikes_count - 1 WHERE id = ?", postId).Error
	return err
}

func (r *PostRepository) Insert(post postDTO.CreatePostDTO) error {
	err := r.db.Exec("INSERT INTO posts (content, title, language_id, address, status, date, views_count, likes_count, dislikes_count, user_id) "+
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", post.Content, post.Title, post.LanguageID, post.Address, postDict.WAITING, time.Now(), 0, 0, 0, post.UserID).Error
	return err
}
