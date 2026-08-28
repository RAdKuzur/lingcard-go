package dictionaryHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"lingcard-go/services/wordTranslationService"
	"net/http"
	"strconv"
)

type DictionaryHandler struct {
	wordTranslationService *wordTranslationService.WordTranslationService
}

func New(wordTranslationService *wordTranslationService.WordTranslationService) *DictionaryHandler {
	return &DictionaryHandler{
		wordTranslationService: wordTranslationService,
	}
}

func (h *DictionaryHandler) Translate(w http.ResponseWriter, r *http.Request) {
	baseLangId, _ := strconv.Atoi(chi.URLParam(r, "baseLanguageId"))
	targetLangId, _ := strconv.Atoi(chi.URLParam(r, "targetLanguageId"))

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	search := r.URL.Query().Get("search")
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	words := h.wordTranslationService.Translate(baseLangId, targetLangId, page, limit, search)
	count := h.wordTranslationService.CountWords(baseLangId, targetLangId, search)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"success":     true,
		"data":        words,
		"amountWords": count,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
