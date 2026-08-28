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
func (r *PostRepository) Find(id int) post.Post {
	var item post.Post
	r.db.Raw("SELECT * FROM posts WHERE id = ?", id).Scan(&item)
	return item
}

func (r *PostRepository) FindApprovedPostsByLanguageId(langId int) []post.Post {
	var posts []post.Post
	r.db.Raw("SELECT * FROM posts WHERE language_id = ? AND status = ? ORDER BY date ASC", langId, postDict.APPROVED).Scan(&posts)
	return posts
}

func (r *PostRepository) IncrementViewsCount(postId int) {
	r.db.Exec("UPDATE posts SET views_count = views_count + 1 WHERE id = ?", postId)
}

func (r *PostRepository) IncrementLikesCount(postId int) {
	r.db.Exec("UPDATE posts SET likes_count = likes_count + 1 WHERE id = ?", postId)
}

func (r *PostRepository) IncrementDislikesCount(postId int) {
	r.db.Exec("UPDATE posts SET dislikes_count = dislikes_count + 1 WHERE id = ?", postId)
}

func (r *PostRepository) DecrementLikesCount(postId int) {
	r.db.Exec("UPDATE posts SET likes_count = likes_count - 1 WHERE id = ?", postId)
}

func (r *PostRepository) DecrementDislikesCount(postId int) {
	r.db.Exec("UPDATE posts SET dislikes_count = dislikes_count - 1 WHERE id = ?", postId)
}

func (r *PostRepository) Insert(post postDTO.CreatePostDTO) {
	r.db.Exec("INSERT INTO posts (content, title, language_id, address, status, date, views_count, likes_count, dislikes_count, user_id) "+
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", post.Content, post.Title, post.LanguageID, post.Address, postDict.WAITING, time.Now(), 0, 0, 0, post.UserID)
}
