package reactionRepository

import (
	"gorm.io/gorm"
	"lingcard-go/dictionaries/reaction"
	reactionDTO "lingcard-go/dto/reaction"
	reactionModel "lingcard-go/models/reaction"
)

type ReactionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ReactionRepository {
	return &ReactionRepository{
		db: db,
	}
}

func (r *ReactionRepository) CountLikes(postId int) int {
	var count int
	r.db.Raw("SELECT COUNT(*) FROM reactions WHERE post_id = ? AND status = ?", postId, reaction.Like).Scan(&count)
	return count
}

func (r *ReactionRepository) CountDislikes(postId int) int {
	var count int
	r.db.Raw("SELECT COUNT(*) FROM reactions WHERE post_id = ? AND status = ?", postId, reaction.Dislike).Scan(&count)
	return count
}

func (r *ReactionRepository) IsLiked(userId, postId int) bool {
	var count int
	r.db.Raw("SELECT COUNT(*) FROM reactions WHERE post_id = ? AND user_id = ? AND status = ?", postId, userId, reaction.Like).Scan(&count)
	return count > 0
}

func (r *ReactionRepository) IsDisliked(userId, postId int) bool {
	var count int
	r.db.Raw("SELECT COUNT(*) FROM reactions WHERE post_id = ? AND user_id = ? AND status = ?", postId, userId, reaction.Dislike).Scan(&count)
	return count > 0
}

func (r *ReactionRepository) DeleteReaction(userId, postId, status int) {
	r.db.Exec("DELETE FROM reactions WHERE post_id = ? AND user_id = ? AND status = ?", postId, userId, status)
}

func (r *ReactionRepository) Delete(id int) {
	r.db.Exec("DELETE FROM reactions WHERE id = ?", id)
}

func (r *ReactionRepository) Insert(dto reactionDTO.ReactionDTO) {
	r.db.Exec("INSERT INTO reactions (user_id, post_id, status) VALUES (?, ?, ?)", dto.UserID, dto.PostID, dto.Status)
}

func (r *ReactionRepository) FindByUserId(userID, postID int) reactionModel.Reaction {
	var item reactionModel.Reaction
	r.db.Raw("SELECT * FROM reactions WHERE post_id = ? AND user_id = ?", postID, userID).Scan(&item)
	return item
}
