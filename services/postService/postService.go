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

func (s *PostService) GetPostsByCode(code string) []post.SimplePostDTO {
	var postsDTO []post.SimplePostDTO

	language := s.languageRepository.FindByCode(code)
	posts := s.postRepository.FindApprovedPostsByLanguageId(language.Id)

	for _, item := range posts {
		likesCount := s.reactionRepository.CountLikes(item.ID)
		dislikesCount := s.reactionRepository.CountDislikes(item.ID)
		User, _ := s.userRepository.One(item.UserID)
		status, _ := postDict.StatusPostDictionary{}.Get(item.Status)
		postsDTO = append(postsDTO, post.SimplePostDTO{
			ID:            item.ID,
			Content:       item.Content,
			Date:          item.Date.Format("2006-01-02 15:04:05"),
			Title:         item.Title,
			Code:          code,
			Username:      User.Name,
			Address:       item.Address,
			Status:        status,
			ViewsCount:    item.ViewsCount,
			LikesCount:    likesCount,
			DislikesCount: dislikesCount,
		})
	}
	return postsDTO
}

func (s *PostService) GetOne(ctx context.Context, id int) post.PostDTO {
	var commentsDTO []post.CommentDTO
	User := ctx.Value("User").(user.User)
	Post := s.postRepository.Find(id)
	isLiked := s.reactionRepository.IsLiked(User.ID, Post.ID)
	isDisliked := s.reactionRepository.IsDisliked(User.ID, Post.ID)
	likesCount := s.reactionRepository.CountLikes(User.ID)
	dislikesCount := s.reactionRepository.CountDislikes(User.ID)
	comments := s.commentRepository.GetAllComments(id)
	for _, comment := range comments {
		User, _ := s.userRepository.One(comment.UserID)
		Language := s.languageRepository.Find(User.BaseLanguageID)
		commentsDTO = append(commentsDTO, post.CommentDTO{
			ID:           comment.ID,
			Text:         comment.Text,
			Username:     User.Name,
			Time:         comment.Time.Format("2006-01-02 15:04:05"),
			IsFixed:      comment.IsFixed,
			LanguageCode: Language.Code,
		})
	}
	s.postRepository.IncrementViewsCount(id)
	status, _ := postDict.StatusPostDictionary{}.Get(Post.Status)
	languagePost := s.languageRepository.Find(Post.LanguageID)
	postDTO := post.PostDTO{
		ID:            Post.ID,
		Content:       Post.Content,
		Date:          Post.Date.Format("2006-01-02 15:04:05"),
		Title:         Post.Title,
		Code:          languagePost.Code,
		Username:      User.Name,
		Address:       Post.Address,
		Status:        status,
		ViewsCount:    Post.ViewsCount,
		LikesCount:    likesCount,
		DislikesCount: dislikesCount,
		IsLiked:       isLiked,
		IsDisliked:    isDisliked,
		Comments:      commentsDTO,
	}
	return postDTO
}

func (s *PostService) CreateComment(ctx context.Context, postId int, text string) {

	User := ctx.Value("User").(user.User)
	var commentDTO = post.CreateCommentDTO{
		PostID: postId,
		UserID: User.ID,
		Text:   text,
	}
	s.commentRepository.Insert(commentDTO)
}

func (s *PostService) Create(ctx context.Context, dto post.CreatePostDTO) {
	User := ctx.Value("User").(user.User)
	dto.UserID = User.ID
	s.postRepository.Insert(dto)
}

func (s *PostService) DeleteComment(ctx context.Context, postId int) {
	User := ctx.Value("User").(user.User)
	Comment := s.commentRepository.Find(postId)
	if User.ID == Comment.UserID {
		s.commentRepository.Delete(postId)
	}
}
