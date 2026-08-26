package languageService

import (
	"lingcard-go/dto/language"
	"lingcard-go/repositories/availableLanguageRepository"
	"lingcard-go/repositories/languageRepository"
)

type LanguageService struct {
	langRepository              *languageRepository.LanguageRepository
	availableLanguageRepository *availableLanguageRepository.AvailableLanguageRepository
}

func New(langRepo *languageRepository.LanguageRepository, availableLangRepo *availableLanguageRepository.AvailableLanguageRepository) *LanguageService {
	return &LanguageService{
		langRepository:              langRepo,
		availableLanguageRepository: availableLangRepo,
	}
}

func (s *LanguageService) All() []language.LangDTO {
	var langDTO []language.LangDTO
	var languages = s.langRepository.All()
	for _, l := range languages {
		var item = language.LangDTO{
			Name: l.Name,
			Code: l.Code,
		}
		langDTO = append(langDTO, item)
	}
	return langDTO
}

func (s *LanguageService) Map() []language.LangMapDTO {
	var langMapDTO []language.LangMapDTO
	var languages = s.langRepository.AllActive()
	for _, l := range languages {
		var availableCodes []string
		for _, baseLanguage := range l.BaseLanguages {
			availableCodes = append(availableCodes, baseLanguage.TargetLanguage.Code)
		}
		langMapDTO = append(langMapDTO, language.LangMapDTO{
			Code:           l.Code,
			Label:          l.Name,
			AvailableCodes: availableCodes,
		})
	}
	return langMapDTO
}

func (s *LanguageService) ExceptLanguage(id int) []language.LangDTO {
	var langDTO []language.LangDTO
	var languages = s.availableLanguageRepository.FindByBaseLanguageId(id)
	for _, l := range languages {
		var item = language.LangDTO{
			ID:   l.TargetLanguage.Id,
			Name: l.TargetLanguage.Name,
			Code: l.TargetLanguage.Code,
		}
		langDTO = append(langDTO, item)
	}
	return langDTO
}

func (s *LanguageService) AllActive() []language.LangDTO {
	var langDTO []language.LangDTO
	var languages = s.langRepository.AllActive()
	for _, l := range languages {
		var item = language.LangDTO{
			ID:   l.Id,
			Name: l.Name,
			Code: l.Code,
		}
		langDTO = append(langDTO, item)
	}
	return langDTO
}
