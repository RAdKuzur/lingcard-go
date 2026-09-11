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

func (s *ReactionService) Like(ctx context.Context, postId int) error {
	User := ctx.Value("User").(user.User)
	var dto = reactionDTO.ReactionDTO{
		UserID: User.ID,
		PostID: postId,
		Status: reaction.LIKE,
	}
	err := s.reactionRepository.DeleteReaction(User.ID, postId, reaction.DISLIKE)
	if err != nil {
		return err
	}
	err1 := s.reactionRepository.Insert(dto)
	if err1 != nil {
		return err1
	}
	err2 := s.postRepository.IncrementLikesCount(postId)
	if err2 != nil {
		return err2
	}
	return nil
}

func (s *ReactionService) Dislike(ctx context.Context, postId int) error {
	usr := ctx.Value("User").(user.User)
	var dto = reactionDTO.ReactionDTO{
		UserID: usr.ID,
		PostID: postId,
		Status: reaction.DISLIKE,
	}
	err := s.reactionRepository.DeleteReaction(usr.ID, postId, reaction.LIKE)
	if err != nil {
		return err
	}
	err1 := s.reactionRepository.Insert(dto)
	if err1 != nil {
		return err1
	}
	err2 := s.postRepository.IncrementDislikesCount(postId)
	if err2 != nil {
		return err2
	}
	return nil
}

func (s *ReactionService) Unset(ctx context.Context, postId int) error {
	usr := ctx.Value("User").(user.User)
	reactionItem, err := s.reactionRepository.FindByUserId(usr.ID, postId)
	if err != nil {
		return err
	}
	err1 := s.reactionRepository.Delete(reactionItem.ID)
	if err1 != nil {
		return err1
	}
	if reactionItem.Status == reaction.LIKE {
		err2 := s.postRepository.DecrementLikesCount(postId)
		if err2 != nil {
			return err2
		}
	}
	if reactionItem.Status == reaction.DISLIKE {
		err3 := s.postRepository.DecrementDislikesCount(postId)
		if err3 != nil {
			return err3
		}
	}
	return nil
}
