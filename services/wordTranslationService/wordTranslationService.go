package wordTranslationService

import (
	"lingcard-go/dictionaries/level"
	"lingcard-go/dto/translation"
	"lingcard-go/repositories/wordRepository"
	"lingcard-go/repositories/wordTranslationRepository"
)

type WordTranslationService struct {
	wordTranslationRepository *wordTranslationRepository.WordTranslationRepository
	wordRepository            *wordRepository.WordRepository
}

func New(wordTranslationRepository *wordTranslationRepository.WordTranslationRepository, wordRepository *wordRepository.WordRepository) *WordTranslationService {
	return &WordTranslationService{
		wordTranslationRepository: wordTranslationRepository,
		wordRepository:            wordRepository,
	}
}

func (s *WordTranslationService) Translate(baseLangId, targetLangId, page, limit int, search string) []translation.WordTranslationDTO {
	var result []translation.WordTranslationDTO
	items := s.wordTranslationRepository.GetPaginateByTargetLanguageIdAndBaseLanguageId(baseLangId, targetLangId, page, limit, search)
	for _, item := range items {
		word := s.wordRepository.Find(item.WordID)
		Level, _ := level.LevelDictionary{}.Get(word.Level)
		result = append(result, translation.WordTranslationDTO{
			ID:            item.ID,
			Translation:   item.Translation,
			Text:          word.Text,
			Transcription: word.Transcription,
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
