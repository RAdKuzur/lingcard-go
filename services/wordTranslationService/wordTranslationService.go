package wordTranslationService

import (
	"lingcard-go/models/words"
	"lingcard-go/repositories/wordTranslationRepository"
)

type WordTranslationService struct {
	wordTranslationRepository *wordTranslationRepository.WordTranslationRepository
}

func New(wordTranslationRepository *wordTranslationRepository.WordTranslationRepository) *WordTranslationService {
	return &WordTranslationService{
		wordTranslationRepository: wordTranslationRepository,
	}
}

func (s *WordTranslationService) Translate(baseLangId, targetLangId, page, limit int, search string) []words.WordTranslationDTO {
	var result []words.WordTranslationDTO
	result = s.wordTranslationRepository.GetPaginateByTargetLanguageIdAndBaseLanguageId(baseLangId, targetLangId, page, limit, search)
	return result
}

func (s *WordTranslationService) CountWords(baseLangId, targetLangId int, search string) int {
	var count int
	count = s.wordTranslationRepository.CountByTargetLanguageIdAndBaseLanguageId(baseLangId, targetLangId, search)
	return count
}
