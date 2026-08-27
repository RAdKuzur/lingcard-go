package wordTranslationService

import (
	"lingcard-go/dictionaries/level"
	"lingcard-go/dto/translation"
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

func (s *WordTranslationService) Translate(baseLangId, targetLangId, page, limit int, search string) []translation.WordTranslationDTO {
	var result []translation.WordTranslationDTO
	items := s.wordTranslationRepository.GetPaginateByTargetLanguageIdAndBaseLanguageId(baseLangId, targetLangId, page, limit, search)
	for _, item := range items {
		Level, _ := level.LevelDictionary{}.Get(item.Level)
		result = append(result, translation.WordTranslationDTO{
			ID:            item.ID,
			Translation:   item.Translation,
			Text:          item.Text,
			Transcription: item.Transcription,
			Level:         Level,
		})
	}
	return result
}

func (s *WordTranslationService) CountWords(baseLangId, targetLangId int, search string) int {
	var count int
	count = s.wordTranslationRepository.CountByTargetLanguageIdAndBaseLanguageId(baseLangId, targetLangId, search)
	return count
}
