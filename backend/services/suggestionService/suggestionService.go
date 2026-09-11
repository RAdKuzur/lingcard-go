package suggestionService

import (
	"context"
	"lingcard-go/dto/suggestion"
	"lingcard-go/models/user"
	"lingcard-go/repositories/suggestionRepository"
)

type SuggestionService struct {
	suggestionRepository *suggestionRepository.SuggestionRepository
}

func New(suggestionRepository *suggestionRepository.SuggestionRepository) *SuggestionService {
	return &SuggestionService{suggestionRepository: suggestionRepository}
}

func (s *SuggestionService) Create(ctx context.Context, dto suggestion.SuggestionDTO) error {
	ctxUser := ctx.Value("User").(user.User)
	dto.UserID = ctxUser.ID
	err := s.suggestionRepository.Insert(dto)
	if err != nil {
		return err
	}
	return nil
}
