package reactionService

import (
	"context"
	"lingcard-go/dictionaries/reaction"
	reactionDTO "lingcard-go/dto/reaction"
	"lingcard-go/models/user"
	"lingcard-go/repositories/postRepository"
	"lingcard-go/repositories/reactionRepository"
)

type ReactionService struct {
	reactionRepository *reactionRepository.ReactionRepository
	postRepository     *postRepository.PostRepository
}

func New(reactionRepository *reactionRepository.ReactionRepository, postRepository *postRepository.PostRepository) *ReactionService {
	return &ReactionService{
		reactionRepository: reactionRepository,
		postRepository:     postRepository,
	}
}

func (s *ReactionService) Like(ctx context.Context, postId int) {
	User := ctx.Value("User").(user.User)
	var dto = reactionDTO.ReactionDTO{
		UserID: User.ID,
		PostID: postId,
		Status: reaction.Like,
	}
	s.reactionRepository.DeleteReaction(User.ID, postId, reaction.Dislike)
	s.reactionRepository.Insert(dto)
	s.postRepository.IncrementLikesCount(postId)

}
func (s *ReactionService) Dislike(ctx context.Context, postId int) {
	User := ctx.Value("User").(user.User)
	var dto = reactionDTO.ReactionDTO{
		UserID: User.ID,
		PostID: postId,
		Status: reaction.Dislike,
	}
	s.reactionRepository.DeleteReaction(User.ID, postId, reaction.Like)
	s.reactionRepository.Insert(dto)
	s.postRepository.IncrementDislikesCount(postId)
}
func (s *ReactionService) Unset(ctx context.Context, postId int) {
	User := ctx.Value("User").(user.User)
	reactionItem := s.reactionRepository.FindByUserId(User.ID, postId)
	s.reactionRepository.Delete(reactionItem.ID)
	if reactionItem.Status == reaction.Like {
		s.postRepository.DecrementLikesCount(postId)
	}
	if reactionItem.Status == reaction.Dislike {
		s.postRepository.DecrementDislikesCount(postId)
	}
}
