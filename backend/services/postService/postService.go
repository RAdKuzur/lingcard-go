package postService

import (
	"context"
	postDict "lingcard-go/dictionaries/post"
	"lingcard-go/dto/post"
	"lingcard-go/models/user"
	"lingcard-go/repositories/commentRepository"
	"lingcard-go/repositories/languageRepository"
	"lingcard-go/repositories/postRepository"
	"lingcard-go/repositories/reactionRepository"
	"lingcard-go/repositories/userRepository"
)

type PostService struct {
	postRepository     *postRepository.PostRepository
	languageRepository *languageRepository.LanguageRepository
	reactionRepository *reactionRepository.ReactionRepository
	commentRepository  *commentRepository.CommentRepository
	userRepository     *userRepository.UserRepository
}

func New(
	postRepository *postRepository.PostRepository,
	languageRepository *languageRepository.LanguageRepository,
	reactionRepository *reactionRepository.ReactionRepository,
	commentRepository *commentRepository.CommentRepository,
	userRepository *userRepository.UserRepository) *PostService {
	return &PostService{
		postRepository:     postRepository,
		languageRepository: languageRepository,
		reactionRepository: reactionRepository,
		commentRepository:  commentRepository,
		userRepository:     userRepository,
	}
}

func (s *PostService) GetPostsByCode(code string) ([]post.SimplePostDTO, error) {
	var postsDTO = make([]post.SimplePostDTO, 0)

	language, err := s.languageRepository.FindByCode(code)
	if err != nil {
		return postsDTO, err
	}
	posts, err2 := s.postRepository.FindApprovedPostsByLanguageId(language.ID)
	if err2 != nil {
		return postsDTO, err2
	}
	for _, item := range posts {
		likesCount, err3 := s.reactionRepository.CountLikes(item.ID)
		if err3 != nil {
			return postsDTO, err3
		}
		dislikesCount, err3 := s.reactionRepository.CountDislikes(item.ID)
		if err3 != nil {
			return postsDTO, err3
		}
		usr, err4 := s.userRepository.One(item.UserID)
		if err4 != nil {
			return postsDTO, err4
		}
		status := postDict.StatusPostDictionary{}.Get(item.Status)
		postsDTO = append(postsDTO, post.SimplePostDTO{
			ID:            item.ID,
			Content:       item.Content,
			Date:          item.Date.Format("2006-01-02 15:04:05"),
			Title:         item.Title,
			Code:          code,
			Username:      usr.Name,
			Address:       item.Address,
			Status:        status,
			ViewsCount:    item.ViewsCount,
			LikesCount:    likesCount,
			DislikesCount: dislikesCount,
		})
	}
	return postsDTO, nil
}

func (s *PostService) GetOne(ctx context.Context, id int) (post.PostDTO, error) {
	commentsDTO := make([]post.CommentDTO, 0)
	usr := ctx.Value("User").(user.User)
	pst, err := s.postRepository.Find(id)
	if err != nil {
		return post.PostDTO{}, err
	}
	pstUser, errUser := s.userRepository.One(pst.UserID)
	if errUser != nil {
		return post.PostDTO{}, errUser
	}
	isLiked, err2 := s.reactionRepository.IsLiked(usr.ID, pst.ID)
	if err2 != nil {
		return post.PostDTO{}, err2
	}
	isDisliked, err3 := s.reactionRepository.IsDisliked(usr.ID, pst.ID)
	if err3 != nil {
		return post.PostDTO{}, err3
	}
	likesCount, err4 := s.reactionRepository.CountLikes(usr.ID)
	if err4 != nil {
		return post.PostDTO{}, err4
	}
	dislikesCount, err5 := s.reactionRepository.CountDislikes(usr.ID)
	if err5 != nil {
		return post.PostDTO{}, err5
	}
	comments, err6 := s.commentRepository.GetAllComments(id)
	if err6 != nil {
		return post.PostDTO{}, err6
	}
	for _, comment := range comments {
		usr2, err7 := s.userRepository.One(comment.UserID)
		if err7 != nil {
			return post.PostDTO{}, err7
		}
		lang, err8 := s.languageRepository.Find(usr2.BaseLanguageID)
		if err8 != nil {
			return post.PostDTO{}, err8
		}
		commentsDTO = append(commentsDTO, post.CommentDTO{
			ID:           comment.ID,
			Text:         comment.Text,
			Username:     usr2.Name,
			Time:         comment.Time.Format("2006-01-02 15:04:05"),
			IsFixed:      comment.IsFixed,
			LanguageCode: lang.Code,
		})
	}
	err9 := s.postRepository.IncrementViewsCount(id)
	if err9 != nil {
		return post.PostDTO{}, err9
	}
	status := postDict.StatusPostDictionary{}.Get(pst.Status)
	languagePost, err10 := s.languageRepository.Find(pst.LanguageID)
	if err10 != nil {
		return post.PostDTO{}, err10
	}
	postDTO := post.PostDTO{
		ID:            pst.ID,
		Content:       pst.Content,
		Date:          pst.Date.Format("2006-01-02 15:04:05"),
		Title:         pst.Title,
		Code:          languagePost.Code,
		Username:      pstUser.Name,
		Address:       pst.Address,
		Status:        status,
		ViewsCount:    pst.ViewsCount,
		LikesCount:    likesCount,
		DislikesCount: dislikesCount,
		IsLiked:       isLiked,
		IsDisliked:    isDisliked,
		Comments:      commentsDTO,
	}
	return postDTO, nil
}

func (s *PostService) CreateComment(ctx context.Context, postId int, text string) error {
	usr := ctx.Value("User").(user.User)
	var commentDTO = post.CreateCommentDTO{
		PostID: postId,
		UserID: usr.ID,
		Text:   text,
	}
	err := s.commentRepository.Insert(commentDTO)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) Create(ctx context.Context, dto post.CreatePostDTO) error {
	usr := ctx.Value("User").(user.User)
	dto.UserID = usr.ID
	err := s.postRepository.Insert(dto)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) DeleteComment(ctx context.Context, postId int) error {
	usr := ctx.Value("User").(user.User)
	comm, err := s.commentRepository.Find(postId)
	if err != nil {
		return err
	}
	if usr.ID == comm.UserID {
		err2 := s.commentRepository.Delete(postId)
		if err2 != nil {
			return err2
		}
	}
	return nil
}
