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

func (s *LanguageService) All() ([]language.LangDTO, error) {
	var langDTO = make([]language.LangDTO, 0)
	languages, err := s.langRepository.All()
	if err != nil {
		return nil, err
	}
	for _, l := range languages {
		var item = language.LangDTO{
			Name: l.Name,
			Code: l.Code,
		}
		langDTO = append(langDTO, item)
	}
	return langDTO, err
}

func (s *LanguageService) Map() ([]language.LangMapDTO, error) {
	var langMapDTO = make([]language.LangMapDTO, 0)
	var languages, err = s.langRepository.AllActive()
	if err != nil {
		return nil, err
	}
	for _, l := range languages {
		var availableCodes []string
		availableBaseLanguages, errBaseLang := s.availableLanguageRepository.FindByBaseLanguageId(l.ID)
		if errBaseLang != nil {
			return nil, errBaseLang
		}
		for _, baseLanguage := range availableBaseLanguages {
			targetLanguage, errTargetLang := s.langRepository.Find(baseLanguage.TargetLanguageID)
			if errTargetLang != nil {
				return nil, errTargetLang
			}
			availableCodes = append(availableCodes, targetLanguage.Code)
		}
		langMapDTO = append(langMapDTO, language.LangMapDTO{
			Code:           l.Code,
			Label:          l.Name,
			AvailableCodes: availableCodes,
		})
	}
	return langMapDTO, err
}

func (s *LanguageService) ExceptLanguage(id int) ([]language.LangDTO, error) {
	var langDTO = make([]language.LangDTO, 0)
	languages, err := s.availableLanguageRepository.FindByBaseLanguageId(id)
	if err != nil {
		return nil, err
	}
	for _, l := range languages {
		lang, err1 := s.langRepository.Find(l.TargetLanguageID)
		if err1 != nil {
			return nil, err1
		}
		var item = language.LangDTO{
			ID:   lang.ID,
			Name: lang.Name,
			Code: lang.Code,
		}
		langDTO = append(langDTO, item)
	}
	return langDTO, err
}

func (s *LanguageService) AllActive() ([]language.LangDTO, error) {
	var langDTO = make([]language.LangDTO, 0)
	var languages, err = s.langRepository.AllActive()
	if err != nil {
		return nil, err
	}
	for _, l := range languages {
		var item = language.LangDTO{
			ID:   l.ID,
			Name: l.Name,
			Code: l.Code,
		}
		langDTO = append(langDTO, item)
	}
	return langDTO, err
}
