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

func (r *ReactionRepository) CountLikes(postId int) (int, error) {
	var count int
	err := r.db.Raw("SELECT COUNT(*) FROM reactions WHERE post_id = ? AND status = ?", postId, reaction.LIKE).Scan(&count).Error
	return count, err
}

func (r *ReactionRepository) CountDislikes(postId int) (int, error) {
	var count int
	err := r.db.Raw("SELECT COUNT(*) FROM reactions WHERE post_id = ? AND status = ?", postId, reaction.DISLIKE).Scan(&count).Error
	return count, err
}

func (r *ReactionRepository) IsLiked(userId, postId int) (bool, error) {
	var count int
	err := r.db.Raw("SELECT COUNT(*) FROM reactions WHERE post_id = ? AND user_id = ? AND status = ?", postId, userId, reaction.LIKE).Scan(&count).Error
	return count > 0, err
}

func (r *ReactionRepository) IsDisliked(userId, postId int) (bool, error) {
	var count int
	err := r.db.Raw("SELECT COUNT(*) FROM reactions WHERE post_id = ? AND user_id = ? AND status = ?", postId, userId, reaction.DISLIKE).Scan(&count).Error
	return count > 0, err
}

func (r *ReactionRepository) DeleteReaction(userId, postId, status int) error {
	err := r.db.Exec("DELETE FROM reactions WHERE post_id = ? AND user_id = ? AND status = ?", postId, userId, status).Error
	return err
}

func (r *ReactionRepository) Delete(id int) error {
	err := r.db.Exec("DELETE FROM reactions WHERE id = ?", id).Error
	return err
}

func (r *ReactionRepository) Insert(dto reactionDTO.ReactionDTO) error {
	err := r.db.Exec("INSERT INTO reactions (user_id, post_id, status) VALUES (?, ?, ?)", dto.UserID, dto.PostID, dto.Status).Error
	return err
}

func (r *ReactionRepository) FindByUserId(userID, postID int) (reactionModel.Reaction, error) {
	var item reactionModel.Reaction
	err := r.db.Raw("SELECT * FROM reactions WHERE post_id = ? AND user_id = ?", postID, userID).Scan(&item).Error
	return item, err
}
