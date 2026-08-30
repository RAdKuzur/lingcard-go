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

func (s *WordTranslationService) Translate(baseLangId, targetLangId, page, limit int, search string) ([]translation.WordTranslationDTO, error) {
	var result = make([]translation.WordTranslationDTO, 0)
	items, err := s.wordTranslationRepository.GetPaginateByTargetLanguageIdAndBaseLanguageId(baseLangId, targetLangId, page, limit, search)
	if err != nil {
		return result, err
	}
	for _, item := range items {
		word, err1 := s.wordRepository.Find(item.WordID)
		if err1 != nil {
			return result, err1
		}
		lvl := level.LevelDictionary{}.Get(word.Level)
		result = append(result, translation.WordTranslationDTO{
			ID:            item.ID,
			Translation:   item.Translation,
			Text:          word.Text,
			Transcription: word.Transcription,
			Level:         lvl,
		})
	}
	return result, err
}

func (s *WordTranslationService) CountWords(baseLangId, targetLangId int, search string) (int, error) {
	count, err := s.wordTranslationRepository.CountByTargetLanguageIdAndBaseLanguageId(baseLangId, targetLangId, search)
	if err != nil {
		return 0, err
	}
	return count, err
}
